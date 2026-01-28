package dbmaster

import (
	"skeleton/infra/ent"
	schemahook "skeleton/infra/ent/schema/hook"
	"skeleton/pkg/config"
)

func RegisterHooks(client *ent.Client, bcryptCfg config.BcryptProperties) {
	if client == nil {
		panic("client cannot be nil")
	}

	client.User.Use(schemahook.NewHashUserPasswordHook(bcryptCfg.Cost))
}
