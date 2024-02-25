package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware function that logs request details.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		fmt.Printf("[%d] %s %s %s %s\n",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.RemoteAddr,
			duration,
		)
	}
}
