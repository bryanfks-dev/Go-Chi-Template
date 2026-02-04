package authusecase

import (
	"context"
	"net/http"
	"skeleton/infra/ent"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *AuthUsecase) ProcessUserLogin(
	ctx context.Context,
	email string,
	passwd string,
	userAgent string,
) (*ent.User, string, string, string, string, error) {
	user, err := u.userRepo.FindByEmail(ctx, u.db, email)
	if err != nil {
		return nil, "", "", "", "", err
	}

	if user == nil {
		u.logger.Debug("user not found", zap.String("email", email))
		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidCredentials,
		)
	}

	if ok := u.sec.CheckPasswordHash(passwd, user.Password); !ok {
		u.logger.Debug("invalid password")
		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidCredentials,
		)
	}

	refreshTokenID, refreshToken, refreshTokenExpiresAt, err :=
		u.sec.NewRefreshJWT(user.ID)
	if err != nil {
		u.logger.Error(
			"failed to generate refresh token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	accessToken, err := u.sec.NewAccessJWT(user.ID, user.Email, "user", nil)
	if err != nil {
		u.logger.Error(
			"failed to generate access token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	xsrfToken, err := uuid.NewV7()
	if err != nil {
		u.logger.Error(
			"failed to generate csrf token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	csrfToken, err := u.sec.NewCSRF(user.ID, xsrfToken.String())
	if err != nil {
		u.logger.Error(
			"failed to generate csrf token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	session := authdomain.NewSessionInput(
		user.ID,
		refreshTokenID,
		refreshToken,
		userAgent,
		*refreshTokenExpiresAt,
	)
	if _, err = u.authRepo.Create(ctx, u.db, session); err != nil {
		return nil, "", "", "", "", err
	}

	return user, refreshToken, accessToken, csrfToken, xsrfToken.String(), nil
}
