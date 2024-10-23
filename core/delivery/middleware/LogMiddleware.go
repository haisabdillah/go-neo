package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/pkg/logging"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // Call the next handler
		latency := time.Since(start).Milliseconds()
		if c.Writer.Status() >= 500 && c.Writer.Status() != 503 {
			logging.Logger.Error("REQUEST",
				"method", c.Request.Method,
				"url", c.Request.URL.String(),
				"status", c.Writer.Status(),
				"ip_address", c.ClientIP(),
				"timestamp", time.Now().UnixNano()/int64(time.Millisecond),
				"latency", latency,
				"auth_id", c.GetUint("authID"),
				"properties", c.GetString("error"),
			)
		} else {
			logging.Logger.Info("REQUEST",
				"method", c.Request.Method,
				"url", c.Request.URL.String(),
				"status", c.Writer.Status(),
				"ip_address", c.ClientIP(),
				"timestamp", time.Now().UnixNano()/int64(time.Millisecond),
				"latency", latency,
				"auth_id", c.GetUint("authID"),
				"properties", c.GetString("error"),
			)
		}

	}

}
