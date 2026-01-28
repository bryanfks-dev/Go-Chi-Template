package security

import (
	"skeleton/infra/ent"

	"github.com/golang-jwt/jwt/v5"
)

var jwtAlgorithm = map[string]jwt.SigningMethod{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
	"RS256": jwt.SigningMethodRS256,
	"RS384": jwt.SigningMethodRS384,
	"RS512": jwt.SigningMethodRS512,
}

func getJWTAlgorithm(algorithm string) jwt.SigningMethod {
	jwt.GetAlgorithms()
	val, ok := jwtAlgorithm[algorithm]
	if !ok {
		return jwt.SigningMethodHS256
	}

	return val
}

func mapUserToUserClaims(user *ent.User) *UserClaims {
	if user == nil {
		return nil
	}

	return &UserClaims{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
