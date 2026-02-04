package authdomain

import (
	"time"
)

type SessionInput struct {
	UserID         int
	RefreshTokenID string
	RefreshToken   string
	UserAgent      string
	ExpiresAt      time.Time
}

func NewSessionInput(
	userID int,
	refreshTokenID string,
	refreshToken string,
	userAgent string,
	expiresAt time.Time,
) *SessionInput {
	return &SessionInput{
		UserID:         userID,
		RefreshTokenID: refreshTokenID,
		RefreshToken:   refreshToken,
		UserAgent:      userAgent,
		ExpiresAt:      expiresAt,
	}
}
