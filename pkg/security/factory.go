package security

import "skeleton/pkg/config"

type Security struct {
	appCfg    *config.ApplicationProperties
	bcryptCfg *config.BcryptProperties
	jwtCfg    *config.JWTProperties
}

func NewSecurity(
	appCfg *config.ApplicationProperties,
	bcryptCfg *config.BcryptProperties,
	jwtCfg *config.JWTProperties,
) *Security {
	return &Security{
		appCfg:    appCfg,
		bcryptCfg: bcryptCfg,
		jwtCfg:    jwtCfg,
	}
}
