package authusecase

import (
	"context"
	"net/http"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"
	"skeleton/pkg/security"
	"strconv"

	"go.uber.org/zap"
)

func (u *AuthUsecase) GenerateAccessToken(
	ctx context.Context,
	refreshClaims *security.JWTClaims,
) (string, error) {
	userID, err := strconv.Atoi(refreshClaims.Subject)
	if err != nil {
		u.logger.Debug(
			"failed to parsed subject from refresh claims",
			zap.String("subject", refreshClaims.Subject),
			zap.Error(err),
		)
		return "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	email, err := u.userRepo.FindEmailByID(ctx, u.db, userID)
	if err != nil {
		return "", err
	}

	accessToken, err := u.sec.NewAccessJWT(email, "user", nil)
	if err != nil {
		return "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrInvalidToken,
		)
	}

	return accessToken, nil
}
