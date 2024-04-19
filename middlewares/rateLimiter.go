package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5) // Rate limit of 1 request per second with a burst of 3 requests

func RateLimitMiddleware(c *gin.Context) {
	if !limiter.Allow() {
		// Exceeded request limit
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
		c.Abort()
		return
	}
	c.Next()
}
