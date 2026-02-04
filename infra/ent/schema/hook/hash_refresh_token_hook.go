package schemahook

import (
	"errors"
	appent "skeleton/infra/ent"
	"skeleton/pkg/security"

	"context"
	"skeleton/infra/ent"
	"skeleton/infra/ent/hook"
)

func NewHashRefreshTokenHook(sec *security.Security) ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(
			func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				aMutation, ok := m.(*appent.AuthSessionMutation)
				if !ok {
					return nil, errors.New("unexpected mutation type")
				}

				if token, ok := aMutation.RefreshToken(); ok {
					hashedToken := sec.HashToken(token)
					aMutation.SetRefreshToken(hashedToken)
				}

				return next.Mutate(ctx, m)
			},
		)
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
