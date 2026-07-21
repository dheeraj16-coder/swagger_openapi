package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// RedisRateLimiter uses Redis for distributed rate limiting
// Works correctly across multiple App Runner instances
type RedisRateLimiter struct {
	client     *redis.Client
	ctx        context.Context
	limit      int // max requests
	windowSecs int // time window in seconds
}

// skipRateLimitPaths are infrastructure endpoints that must NEVER be rate
// limited: they're polled by trusted internal systems (Prometheus/ADOT
// scraping /metrics, load balancers and AWS health checks hitting /health)
// on a tight timer. Rate limiting them (a) wastes a Redis INCR+EXPIRE on
// every poll, and (b) eventually 429s the poller, breaking metrics and
// health reporting. These must short-circuit BEFORE the Redis call.
var skipRateLimitPaths = map[string]bool{
	"/metrics": true,
	"/health":  true,
}

// NewRedisRateLimiter creates a Redis-backed rate limiter
// limit: max requests per window
// windowSecs: time window in seconds
func NewRedisRateLimiter(limit, windowSecs int) *RedisRateLimiter {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://redis:6379"
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(fmt.Sprintf("invalid REDIS_URL for rate limiter: %v", err))
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("rate limiter cannot connect to Redis: %v", err))
	}

	return &RedisRateLimiter{
		client:     client,
		ctx:        ctx,
		limit:      limit,
		windowSecs: windowSecs,
	}
}

// Allow checks if the IP is allowed to make a request
// Uses Redis INCR + EXPIRE for atomic sliding window
func (r *RedisRateLimiter) Allow(ip string) (bool, int) {
	key := fmt.Sprintf("ratelimit:%s", ip)

	// Atomic increment
	count, err := r.client.Incr(r.ctx, key).Result()
	if err != nil {
		// If Redis fails, allow request (fail open)
		return true, 0
	}

	// Set expiry only on first request in window
	if count == 1 {
		r.client.Expire(r.ctx, key, time.Duration(r.windowSecs)*time.Second)
	}

	remaining := r.limit - int(count)
	if remaining < 0 {
		remaining = 0
	}

	return count <= int64(r.limit), remaining
}

// RateLimitMiddleware creates a Gin middleware using Redis
func RateLimitMiddleware(limiter *RedisRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Infrastructure endpoints bypass rate limiting entirely — no Redis
		// touch, no counter, no 429. Checked before anything else so internal
		// pollers never consume quota or get blocked.
		if skipRateLimitPaths[c.FullPath()] {
			c.Next()
			return
		}

		ip := c.ClientIP()

		allowed, remaining := limiter.Allow(ip)

		// Add rate limit headers
		c.Header("X-RateLimit-Limit", strconv.Itoa(limiter.limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("X-RateLimit-Window", strconv.Itoa(limiter.windowSecs)+"s")

		if !allowed {
			RecordRateLimitBlock(ip, c.FullPath())
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":     "Rate limit exceeded",
				"message":   "Too many requests. Please try again later.",
				"ip":        ip,
				"limit":     limiter.limit,
				"window":    fmt.Sprintf("%ds", limiter.windowSecs),
				"remaining": 0,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
