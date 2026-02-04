package authusecase

import (
	"context"
	"net/http"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"

	"go.uber.org/zap"
)

func (u *AuthUsecase) ProcessUserLogout(
	ctx context.Context,
	refreshToken string,
) error {
	refreshClaims, err := u.decodeRefreshToken(refreshToken)
	if err != nil {
		u.logger.Error("failed to decode refresh token", zap.Error(err))
		return apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	session, err := u.authRepo.FindByRefreshTokenID(
		ctx,
		u.db,
		refreshClaims.ID,
	)
	if err != nil {
		return err
	}

	if session == nil {
		u.logger.Debug("no session found for given refresh token ID")
		return apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	if !u.sec.CompareHashAndToken(session.RefreshToken, refreshToken) {
		u.logger.Debug("refresh token hash does not match stored hash")
		return apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	if err = u.authRepo.DeleteByID(ctx, u.db, session.ID); err != nil {
		return err
	}

	return nil
}
