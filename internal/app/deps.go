package app

import (
	"wallet/internal/shared/config"

	"go.uber.org/zap"
)

type Deps struct {
	Cfg *config.Config
	Log *zap.Logger
}
