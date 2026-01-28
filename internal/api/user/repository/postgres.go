package userrepository

import (
	"context"
	"skeleton/infra/ent"

	"go.uber.org/zap"
)

func (r *UserRepository) FindByID(
	ctx context.Context,
	tx *ent.Client,
	id int,
) (*ent.User, error) {
	u, err := tx.User.Get(ctx, id)
	if err != nil {
		r.logger.Error(
			"Failed finding user by ID",
			zap.Error(err),
			zap.Int("id", id),
		)
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindAll() {
	// Implement your logic here
}
