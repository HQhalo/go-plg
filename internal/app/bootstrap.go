package app

import (
	bypurchase "wallet/internal/earn/by_purchase"
	"wallet/internal/shared/config"
	"wallet/internal/shared/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BootstrapResult struct {
	Engine *gin.Engine
	Cfg    *config.Config
	Log    *zap.Logger
}

func Bootstrap() (*BootstrapResult, error) {
	deps, err := initDeps()
	if err != nil {
		return nil, err
	}

	if deps.Cfg.App.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(
		gin.Recovery(),
		logger.ZapLogger(deps.Log, deps.Cfg.App.Env),
	)

	v1 := r.Group("/v1")

	bypurchase.RegisterRoutes(v1, deps.Log)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &BootstrapResult{
		Engine: r,
		Cfg:    deps.Cfg,
		Log:    deps.Log,
	}, nil
}

func initDeps() (*Deps, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	log, err := logger.NewLogger(cfg.App.Env)
	if err != nil {
		return nil, err
	}

	return &Deps{
		Cfg: cfg,
		Log: log,
	}, nil
}
