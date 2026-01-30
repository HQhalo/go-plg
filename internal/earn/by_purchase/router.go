package bypurchase

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterRoutes(r gin.IRoutes, log *zap.Logger) {
	handler := NewHandler(log)

	r.POST("/earn/purchase", handler.EarnByPurchase)
}
