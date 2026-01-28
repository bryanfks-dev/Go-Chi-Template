ATLAS_PATH = atlas
MIGRATION_PATH = file://infra/db/migrations
ENT_SCHEMA_PATH = ent://infra/ent/schema

DATABASE_PROVIDER = postgres
DATABASE_USER = postgres
DATABASE_PASSWORD =
DATABASE_HOST = 192.168.18.224
DATABASE_PORT = 5432
DATABASE_NAME = chi
DATABASE_SSLMODE = disable

DATABASE_DSN = $(DATABASE_PROVIDER)://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSLMODE)

install:
	go mod tidy
	go mod download

generate: install
	go generate ./...

docs: install
	go tool swag init -d ./ -g ./cmd/http/main.go --parseDependency

format: install
	gofmt -w . && go tool golines -m 80 -w .

start: install
	go tool air

test: install
	go test -v ./...

migrate_create: install
ifeq ($(strip $(name)),)
	echo "Error: name variable is not set. Usage: make create_migration name=<migration_name>"
	exit 1
endif
	$(ATLAS_PATH) migrate diff $(name) \
    	--dir $(MIGRATION_PATH) \
    	--to $(ENT_SCHEMA_PATH) \
    	--dev-url $(DATABASE_DSN)

migrate_apply: install
	$(ATLAS_PATH) migrate apply \
		--dir $(MIGRATION_PATH) \
		--url $(DATABASE_DSN)
