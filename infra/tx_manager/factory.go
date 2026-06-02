package txmanager

import (
	"context"
	"skeleton/infra/ent"
	"skeleton/pkg/consts"
	"skeleton/pkg/logger"

	"go.uber.org/zap"
)

type TxManager struct {
	client *ent.Client
	logger *logger.Logger
}

func NewTxManager(client *ent.Client, logger *logger.Logger) *TxManager {
	return &TxManager{
		client: client,
		logger: logger,
	}
}

func (m *TxManager) WithTx(
	ctx context.Context,
	domainErr error,
	fn func(txCtx context.Context) error,
) error {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		m.logger.Error("Failed to start transaction", zap.Error(err))
		return domainErr
	}

	txCtx := context.WithValue(ctx, consts.DatabaseTranscationContextKey, tx)

	defer func() {
		if pan := recover(); pan != nil {
			m.logger.Error("Panic called in transaction", zap.Any("panic", pan))
			if err := tx.Rollback(); err != nil {
				m.logger.Error("Failed to rollback transaction", zap.Error(err))
			}
			panic(pan)
		}
	}()

	if err := fn(txCtx); err != nil {
		if err := tx.Rollback(); err != nil {
			m.logger.Error("Failed to rollback transaction", zap.Error(err))
			return domainErr
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		m.logger.Error("Failed to commit transaction", zap.Error(err))
		return domainErr
	}
	return nil
}
