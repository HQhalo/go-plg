package logger

import (
	"time"
	"wallet/internal/shared/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewLogger(env string) (*zap.Logger, error) {
	switch env {
	case "production":
		return zap.NewProduction()
	default:
		return zap.NewDevelopment()
	}
}

func ZapLogger(log *zap.Logger, env string) gin.HandlerFunc {
	isProd := env == config.ENV_PRODUCTION

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()
		status := c.Writer.Status()
		if isProd && status < 500 {
			return
		}

		log.Info("http request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}
