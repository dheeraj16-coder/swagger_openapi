package openapi

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	ttl    time.Duration
	ctx    context.Context
}

func NewCache(ttl time.Duration) *Cache {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://redis:6379"
	}

	ctx := context.Background()

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("[CACHE] Invalid REDIS_URL: %v — running without cache", err)
		return &Cache{client: nil, ttl: ttl, ctx: ctx}
	}

	client := redis.NewClient(opt)

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("[CACHE] Warning: Redis not reachable: %v — running without cache", err)
		return &Cache{client: nil, ttl: ttl, ctx: ctx}
	}

	log.Printf("[CACHE] Connected to Redis")
	return &Cache{client: client, ttl: ttl, ctx: ctx}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if c.client == nil {
		return nil, false
	}
	val, err := c.client.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	}
	if err != nil {
		log.Printf("[CACHE] Get error for key %s: %v", key, err)
		return nil, false
	}
	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		log.Printf("[CACHE] Unmarshal error for key %s: %v", key, err)
		return nil, false
	}
	return result, true
}

func (c *Cache) Set(key string, value interface{}) {
	if c.client == nil {
		return
	}
	data, err := json.Marshal(value)
	if err != nil {
		log.Printf("[CACHE] Marshal error for key %s: %v", key, err)
		return
	}
	if err := c.client.Set(c.ctx, key, data, c.ttl).Err(); err != nil {
		log.Printf("[CACHE] Set error for key %s: %v", key, err)
	}
}
