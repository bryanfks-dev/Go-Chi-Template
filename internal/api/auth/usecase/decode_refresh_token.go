package authusecase

import (
	"net/http"
	authdomain "skeleton/internal/api/auth/domain"
	apperror "skeleton/pkg/error"
	"skeleton/pkg/security"
	"strings"

	"go.uber.org/zap"
)

func (u *AuthUsecase) decodeRefreshToken(
	refreshToken string,
) (*security.JWTClaims, error) {
	if strings.TrimSpace(refreshToken) == "" {
		u.logger.Debug("empty refresh token")
		return nil, apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	refreshClaims, err := u.sec.DecodeJWT(refreshToken)
	if err != nil {
		u.logger.Error("failed to decode refresh token", zap.Error(err))
		return nil, apperror.NewAppError(
			http.StatusUnauthorized,
			authdomain.ErrInvalidToken,
		)
	}

	return refreshClaims, nil
}
