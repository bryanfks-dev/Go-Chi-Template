package utils

import (
	"context"
	"skeleton/infra/ent"
	"skeleton/pkg/consts"
	apperror "skeleton/pkg/error"
)

func GetUserIDFromContext(ctx context.Context) (int, error) {
	id, ok := ctx.Value(consts.UserIDContextKey).(int)
	if !ok {
		return 0, apperror.ErrAuthSessionError
	}
	return id, nil
}

func GetDatabaseTransactionFromContext(ctx context.Context) *ent.Tx {
	tx, _ := ctx.Value(consts.DatabaseTranscationContextKey).(*ent.Tx)
	return tx
}
