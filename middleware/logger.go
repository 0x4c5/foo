package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		cost := time.Since(start)
		zap.L().Info(fmt.Sprintf("%d", ctx.Writer.Status()),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("method", ctx.Request.Method),
			zap.Duration("cost", cost),
		)
	}
}
