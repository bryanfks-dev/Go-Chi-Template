package schemahook

import (
	"errors"
	appent "skeleton/infra/ent"
	"skeleton/pkg/security"

	"context"
	"skeleton/infra/ent"
	"skeleton/infra/ent/hook"
)

func NewHashUserPasswordHook(hashCost int) ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(
			func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				uMutation, ok := m.(*appent.UserMutation)
				if !ok {
					return nil, errors.New("Unexpected mutation type")
				}

				if pwd, ok := uMutation.Password(); ok {
					hashedPwd, err := security.HashPassword(pwd, hashCost)
					if err != nil {
						return nil, err
					}
					uMutation.SetPassword(hashedPwd)
				}

				return next.Mutate(ctx, m)
			},
		)
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
