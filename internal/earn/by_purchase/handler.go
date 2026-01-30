package bypurchase

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	log *zap.Logger
}

func NewHandler(log *zap.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

// POST /v1/earn/purchase
func (h *Handler) EarnByPurchase(c *gin.Context) {
	h.log.Info("EarnByPurchase called")
	c.JSON(200, gin.H{
		"message": "Earned by purchase",
	})
}
