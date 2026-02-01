package db

import (
	"context"
	"fmt"
	"time"

	"wallet/internal/shared/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, cfg *config.Config) (*DB, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	poolCfg.MaxConns = int32(cfg.DB.Maxconns)
	poolCfg.MinConns = int32(cfg.DB.Minconns)

	poolCfg.MaxConnIdleTime = 5 * time.Minute
	poolCfg.MaxConnLifetime = 30 * time.Minute
	poolCfg.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{Pool: pool}, nil
}

func (d *DB) Close() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}

func (d *DB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return d.Pool.Begin(ctx)
}
