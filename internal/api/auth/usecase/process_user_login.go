package authusecase

import (
	"context"
	"net/http"
	"skeleton/infra/ent"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"

	"go.uber.org/zap"
)

func (u *AuthUsecase) ProcessUserLogin(
	ctx context.Context,
	email string,
	passwd string,
) (*ent.User, string, string, error) {
	user, err := u.userRepo.FindByEmail(ctx, u.db, email)
	if err != nil {
		return nil, "", "", err
	}

	if user == nil {
		u.logger.Debug("user not found", zap.String("email", email))
		return nil, "", "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidCredentials,
		)
	}

	u.logger.Debug("user found", zap.Int("userID", user.ID))
	if ok := u.sec.CheckPasswordHash(passwd, user.Password); !ok {
		u.logger.Debug("invalid password")
		return nil, "", "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidCredentials,
		)
	}

	refreshToken, err := u.sec.NewRefreshJWT(user.ID)
	if err != nil {
		u.logger.Error(
			"failed to generate refresh token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	accessToken, err := u.sec.NewAccessJWT(user.Email, "user", nil)
	if err != nil {
		u.logger.Error(
			"failed to generate access token",
			zap.Int("userID", user.ID),
			zap.Error(err),
		)

		return nil, "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrGenerateTokenFailed,
		)
	}

	return user, refreshToken, accessToken, nil
}
