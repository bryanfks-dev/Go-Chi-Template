package security

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	Type        string `json:"typ"`
	*UserClaims `json:"omitempty"`
	jwt.RegisteredClaims
}

type UserClaims struct {
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

func (s *Security) NewRefreshJWT(
	userID int,
) (string, error) {
	claimsID, _ := uuid.NewV7()
	claims := JWTClaims{
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       claimsID.String(),
			Subject:  strconv.Itoa(userID),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   s.appCfg.Name,
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					time.Minute *
						time.Duration(s.jwtCfg.Refresh.ExpirationMinutes),
				),
			),
		},
	}

	hashAlgo := getJWTAlgorithm(s.jwtCfg.Algorithm)
	token, err := jwt.
		NewWithClaims(hashAlgo, claims).
		SignedString([]byte(s.jwtCfg.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Security) NewAccessJWT(
	userEmail string,
	userRole string,
	userPermissions *[]string,
) (string, error) {
	if userPermissions == nil {
		userPermissions = &[]string{"*"}
	}

	claimsID, _ := uuid.NewV7()
	claims := JWTClaims{
		Type: "access",
		UserClaims: &UserClaims{
			Email:       userEmail,
			Role:        userRole,
			Permissions: *userPermissions,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       claimsID.String(),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   s.appCfg.Name,
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					time.Minute *
						time.Duration(s.jwtCfg.Access.ExpirationMinutes),
				),
			),
		},
	}

	hashAlgo := getJWTAlgorithm(s.jwtCfg.Algorithm)
	token, err := jwt.
		NewWithClaims(hashAlgo, claims).
		SignedString([]byte(s.jwtCfg.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Security) DecodeJWT(
	token string,
) (*JWTClaims, error) {
	hashAlgo := getJWTAlgorithm(s.jwtCfg.Algorithm)
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != hashAlgo {
				return nil, jwt.ErrTokenUnverifiable
			}
			return []byte(s.jwtCfg.Secret), nil
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

func getJWTAlgorithm(algorithm string) jwt.SigningMethod {
	val, ok := JWTAlgorithm[algorithm]
	if !ok {
		return jwt.SigningMethodHS256
	}
	return val
}
