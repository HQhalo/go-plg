package record

import (
	"context"
	"wallet/internal/shared/db/sqlc"
	"wallet/internal/shared/tx"

	"go.uber.org/zap"
)

type Service struct {
	log *zap.Logger
	tx  *tx.Manager
}

func NewService(log *zap.Logger, txManager *tx.Manager) *Service {
	return &Service{
		log: log,
		tx:  txManager,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, cmd sqlc.CreateLedgerEntryParams) error {
	return s.tx.WithTx(ctx, func(exec tx.Executor) error {
		repo := sqlc.New(exec)
		if _, err := repo.CreateLedgerEntry(ctx, cmd); err != nil {
			return err
		}
		return nil
	})
}
