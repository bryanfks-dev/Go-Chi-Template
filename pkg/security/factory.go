package security

import "skeleton/pkg/config"

type Security struct {
	appCfg    *config.ApplicationProperties
	bcryptCfg *config.BcryptProperties
	hmacCfg   *config.HMACProperties
	jwtCfg    *config.JWTProperties
}

func NewSecurity(
	appCfg *config.ApplicationProperties,
	bcryptCfg *config.BcryptProperties,
	hmacCfg *config.HMACProperties,
	jwtCfg *config.JWTProperties,
) *Security {
	return &Security{
		appCfg:    appCfg,
		bcryptCfg: bcryptCfg,
		hmacCfg:   hmacCfg,
		jwtCfg:    jwtCfg,
	}
}
