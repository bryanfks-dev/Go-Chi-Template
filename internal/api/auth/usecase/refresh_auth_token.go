package authusecase

import (
	"context"
	"net/http"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"
	"strconv"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *AuthUsecase) RefreshAuthToken(
	ctx context.Context,
	refreshToken string,
) (string, string, string, error) {
	refreshClaims, err := u.decodeRefreshToken(refreshToken)
	if err != nil {
		return "", "", "", err
	}

	userID, err := strconv.Atoi(refreshClaims.Subject)
	if err != nil {
		u.logger.Debug(
			"failed to parsed subject from refresh claims",
			zap.String("subject", refreshClaims.Subject),
			zap.Error(err),
		)
		return "", "", "", apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	email, err := u.userRepo.FindEmailByID(ctx, u.db, userID)
	if err != nil {
		return "", "", "", err
	}

	accessToken, err := u.sec.NewAccessJWT(userID, email, "user", nil)
	if err != nil {
		return "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrInvalidToken,
		)
	}

	xsrfToken, err := uuid.NewV7()
	if err != nil {
		u.logger.Error("failed to generate xsrf token", zap.Error(err))
		return "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrInvalidToken,
		)
	}

	csrfToken, err := u.sec.NewCSRF(userID, xsrfToken.String())
	if err != nil {
		u.logger.Error("failed to generate csrf token", zap.Error(err))
		return "", "", "", apperror.NewAppError(
			http.StatusInternalServerError,
			authdomain.ErrInvalidToken,
		)
	}

	return accessToken, csrfToken, xsrfToken.String(), nil
}
