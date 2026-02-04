ATLAS_PATH = atlas
MIGRATION_PATH = file://infra/db/migrations
ENT_SCHEMA_PATH = ent://infra/ent/schema

MIGRATION_DATABASE_PROVIDER = postgres
MIGRATION_DATABASE_USER = postgres
MIGRATION_DATABASE_PASSWORD =
MIGRATION_DATABASE_HOST = 192.168.18.224
MIGRATION_DATABASE_PORT = 5432
MIGRATION_DATABASE_NAME = skeleton
MIGRATION_DATABASE_SSLMODE = disable

MIGRATION_DATABASE_DSN = $(MIGRATION_DATABASE_PROVIDER)://$(MIGRATION_DATABASE_USER):$(MIGRATION_DATABASE_PASSWORD)@$(MIGRATION_DATABASE_HOST):$(MIGRATION_DATABASE_PORT)/$(MIGRATION_DATABASE_NAME)?sslmode=$(MIGRATION_DATABASE_SSLMODE)

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

schema_create: install
ifeq ($(strip $(name)),)
		echo "Error: name variable is not set. Usage: make schema_create name=<schema_name>"
		exit 1
endif
	go run entgo.io/ent/cmd/ent@latest --target infra/ent/schema new $(name)

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

docs_generate: install
	go tool swag init -d ./ -g cmd/http/main.go --parseDependency

migration_status: install
	$(ATLAS_PATH) migrate status \
		--dir $(MIGRATION_PATH) \
		--url $(DATABASE_DSN)

migration_generate: install
ifeq ($(strip $(name)),)
	echo "Error: name variable is not set. Usage: make migration_generate name=<migration_name>"
	exit 1
endif
	$(ATLAS_PATH) migrate diff $(name) \
    	--dir $(MIGRATION_PATH) \
    	--to $(ENT_SCHEMA_PATH) \
    	--dev-url $(MIGRATION_DATABASE_DSN)

migration_hash: install
	$(ATLAS_PATH) migrate hash \
		--dir $(MIGRATION_PATH)

migration_apply: install
	$(ATLAS_PATH) migrate apply \
		--dir $(MIGRATION_PATH) \
		--url $(DATABASE_DSN)
