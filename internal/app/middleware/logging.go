package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func WithLogging(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		logger.Infow("HTTP request log",
			"method", c.Request.Method,
			"uri", c.Request.RequestURI,
			"status", c.Writer.Status(),
			"duration", duration,
			"size", c.Writer.Size(),
		)
	}
}
