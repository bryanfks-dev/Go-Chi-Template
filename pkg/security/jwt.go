package security

import (
	"skeleton/infra/ent"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Type string      `json:"type"`
	User *UserClaims `json:"user,omitempty"`
	jwt.RegisteredClaims
}

type UserClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewJWT(
	jwtType string,
	secret string,
	algorithm string,
	expiredAtMinutes int,
	user *ent.User,
) (*string, error) {
	if user == nil {
		panic("user cannot be nil")
	}

	userClaims := mapUserToUserClaims(user)
	claims := JWTClaims{
		Type: jwtType,
		User: userClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.FullName(),
			ID:      strconv.Itoa(user.ID),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Minute * time.Duration(expiredAtMinutes)),
			),
		},
	}

	jwtAlgorithm := getJWTAlgorithm(algorithm)
	token := jwt.NewWithClaims(jwtAlgorithm, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func DecodeJWT(
	token string,
	secret string,
	algorithm string,
) (*JWTClaims, error) {
	jwtAlgorithm := getJWTAlgorithm(algorithm)
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != jwtAlgorithm {
				return nil, jwt.ErrTokenUnverifiable
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
