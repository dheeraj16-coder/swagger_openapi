package main

import (
	"log"
	sw "my-go-backend/go"
	"my-go-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	routes := sw.ApiHandleFunctions{}

	log.Printf("Server started with rate limiting and Prometheus metrics")

	router := gin.Default()

	
	rateLimiter := middleware.NewRedisRateLimiter(100, 60)

	
	router.Use(middleware.PrometheusMiddleware())
	router.Use(middleware.RateLimitMiddleware(rateLimiter))

	
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	sw.NewRouterWithGinEngine(router, routes)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Printf("Rate limiting: 100 req/60s per IP (Redis distributed)")
	log.Printf("Metrics available at: http://localhost:8080/metrics")
	log.Printf("Listening on :8080")
	log.Fatal(router.Run(":8080"))
}
