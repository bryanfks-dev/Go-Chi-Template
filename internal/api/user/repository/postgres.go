package userrepository

import (
	"context"
	"net/http"
	"skeleton/infra/ent"
	"skeleton/infra/ent/user"
	userdomain "skeleton/internal/api/user/domain"
	apperror "skeleton/pkg/error"

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
			"failed finding user by ID",
			zap.Error(err),
			zap.Int("id", id),
		)
		return nil, apperror.NewAppError(
			http.StatusInternalServerError,
			userdomain.ErrGetUserFailed,
		)
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	tx *ent.Client,
	email string,
) (*ent.User, error) {
	u, err := tx.User.
		Query().
		Where(user.EmailEQ(email)).
		First(ctx)

	if ent.IsNotFound(err) {
		r.logger.Debug("user not found by email", zap.String("email", email))
		return nil, nil
	}
	if err != nil {
		r.logger.Error(
			"failed finding user by email",
			zap.Error(err),
			zap.String("email", email),
		)
		return nil, apperror.NewAppError(
			http.StatusInternalServerError,
			userdomain.ErrGetUserFailed,
		)
	}
	return u, nil
}

func (r *UserRepository) FindEmailByID(
	ctx context.Context,
	tx *ent.Client,
	id int,
) (string, error) {
	e, err := tx.User.
		Query().
		Where(user.ID(id)).
		Select(user.FieldEmail).
		String(ctx)

	if err != nil {
		r.logger.Error(
			"failed finding user email by ID",
			zap.Error(err),
			zap.Int("id", id),
		)
		return "", apperror.NewAppError(
			http.StatusInternalServerError,
			userdomain.ErrGetUserFailed,
		)
	}
	return e, nil
}

func (r *UserRepository) FindAll() {
	// Implement your logic here
}
