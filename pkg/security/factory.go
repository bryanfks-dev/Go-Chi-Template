package security

import "skeleton/pkg/config"

type Security struct {
	bcryptCost int
}

func NewSecurity(cfg *config.BcryptProperties) *Security {
	return &Security{
		bcryptCost: cfg.Cost,
	}
}
