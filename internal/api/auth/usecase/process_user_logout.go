package authusecase

import (
	"context"
	"net/http"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"
)

func (u *AuthUsecase) ProcessUserLogout(
	ctx context.Context,
	refreshToken string,
) error {
	session, err := u.authRepo.FindByRefreshToken(ctx, u.db, refreshToken)
	if err != nil {
		return err
	}

	if session == nil {
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
