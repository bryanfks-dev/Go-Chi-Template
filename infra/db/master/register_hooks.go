package dbmaster

import (
	"skeleton/infra/ent"
	schemahook "skeleton/infra/ent/schema/hook"
	"skeleton/pkg/security"
)

func RegisterHooks(client *ent.Client, security *security.Security) {
	if client == nil {
		panic("client cannot be nil")
	}

	client.User.Use(schemahook.NewHashUserPasswordHook(security))
}
