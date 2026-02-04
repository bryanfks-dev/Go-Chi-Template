package security

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/golang-jwt/jwt/v5"
)

var JWTAlgorithms = map[string]jwt.SigningMethod{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
	"RS256": jwt.SigningMethodRS256,
	"RS384": jwt.SigningMethodRS384,
	"RS512": jwt.SigningMethodRS512,
}

var HMACAlgorithms = map[string]func() hash.Hash{
	"SHA256": sha256.New,
	"SHA512": sha512.New,
}
