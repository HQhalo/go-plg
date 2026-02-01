package create

import (
	"wallet/internal/shared/tx"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterRoutes(r gin.IRoutes, log *zap.Logger, txManager *tx.Manager) {
	service := NewService(log, txManager)
	handler := NewHandler(log, service)

	r.POST("/transaction/create", handler.CreateTransaction)
}
