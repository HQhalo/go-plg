package app

import (
	"context"
	"time"
	"wallet/internal/shared/config"
	"wallet/internal/shared/db"
	"wallet/internal/shared/logger"
	"wallet/internal/shared/tx"
	trxCreate "wallet/internal/transaction/create"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BootstrapResult struct {
	Engine *gin.Engine
	Cfg    *config.Config
	Log    *zap.Logger
}

type Deps struct {
	Cfg *config.Config
	Log *zap.Logger
	DB  *db.DB
	Tx  *tx.Manager
}

func Bootstrap(ctx context.Context) (*BootstrapResult, error) {
	deps, err := initDeps(ctx)
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

	trxCreate.RegisterRoutes(v1, deps.Log, deps.Tx)

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

func initDeps(ctx context.Context) (*Deps, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	log, err := logger.NewLogger(cfg.App.Env)
	if err != nil {
		return nil, err
	}

	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	db, err := db.NewDB(dbCtx, cfg)
	if err != nil {
		return nil, err
	}

	txManager := tx.NewManager(db.Pool)

	return &Deps{
		Cfg: cfg,
		Log: log,
		DB:  db,
		Tx:  txManager,
	}, nil
}
