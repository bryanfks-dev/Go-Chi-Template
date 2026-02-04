package authrepository

import (
	"context"
	"net/http"
	"skeleton/infra/ent"
	"skeleton/infra/ent/authsession"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"

	"go.uber.org/zap"
)

func (r *AuthRepository) Create(
	ctx context.Context,
	tx *ent.Client,
	ss *authdomain.SessionInput,
) (*ent.AuthSession, error) {
	session, err := tx.AuthSession.
		Create().
		SetUserID(ss.UserID).
		SetRefreshTokenID(ss.RefreshTokenID).
		SetRefreshToken(ss.RefreshToken).
		SetUserAgent(ss.UserAgent).
		SetExpiresAt(ss.ExpiresAt).
		Save(ctx)

	if err != nil {
		r.logger.Error(
			"failed creating auth session",
			zap.Int("userID", ss.UserID),
			zap.Error(err),
		)
		return nil, apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrCreateAuthSessionFailed,
		)
	}
	return session, nil
}

func (r *AuthRepository) FindByRefreshTokenID(
	ctx context.Context,
	tx *ent.Client,
	refreshTokenID string,
) (*ent.AuthSession, error) {
	s, err := tx.AuthSession.
		Query().
		Where(authsession.RefreshTokenIDEQ(refreshTokenID)).
		Only(ctx)

	if ent.IsNotFound(err) {
		r.logger.Debug(
			"auth session not found by refresh token ID",
			zap.String("refreshTokenID", refreshTokenID),
		)
		return nil, nil
	}
	if err != nil {
		r.logger.Error(
			"failed finding auth session by refresh token ID",
			zap.Error(err),
		)
		return nil, apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGetAuthSessionFailed,
		)
	}
	return s, nil
}

func (r *AuthRepository) DeleteByID(
	ctx context.Context,
	tx *ent.Client,
	id int,
) error {
	err := tx.AuthSession.DeleteOneID(id).Exec(ctx)
	if err != nil {
		r.logger.Error(
			"failed deleting auth session by ID",
			zap.Int("id", id),
			zap.Error(err),
		)
		return apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrDeleteAuthSessionFailed,
		)
	}
	return nil
}
