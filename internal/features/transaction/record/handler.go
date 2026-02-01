package record

import (
	"wallet/internal/shared/db/sqlc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	log     *zap.Logger
	service *Service
}

func NewHandler(log *zap.Logger, service *Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

// POST /v1/transaction/create
func (h *Handler) CreateTransaction(c *gin.Context) {
	err := h.service.CreateTransaction(c.Request.Context(), sqlc.CreateLedgerEntryParams{})
	if err != nil {
		h.log.Error("failed to create transaction", zap.Error(err))
		c.JSON(500, gin.H{
			"error": "failed to create transaction",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "transaction created",
	})
}
