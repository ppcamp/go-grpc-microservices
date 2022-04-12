default: help

include .env

# COLORS
COLOR_BASE=\033[0m
COLOR_DIM=\e[2m
COLOR_CMD=\x1b[36m
COLOR_DESCR=\x1b[0m
N?=1

.PHONY: run
run: ## Run the server
	@cd src/ && go run cmd/*.go


.PHONY: build
build: ## Build the server
	@cd src/ && go build -race cmd/*.go


.PHONY: migrate
migrate: ## Run migrations
	@echo "Running migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose up


.PHONY: create_migration
create_migration: ## Create a new migration, e.g name=teste make create_migration
	@echo "Creating migration"
	@migrate create -ext sql -dir migrations -seq ${name}


.PHONY: revert_migrations
revert_migrations: ## Revert a given migration, e.g N=2 make revert_migrations, by default 1
	@echo "Reverting migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose down ${N}


.PHONY: generate
generate: ## Create golang protos
	@echo "Generating protos"
	@echo " - Generating messages"
	@protoc --go_out=. protos/*.proto
	@echo " - Generating services"
	@protoc --go-grpc_out=. protos/*.proto


.PHONY: setup_dev
setup_dev: ## Install dev dependencies
	@echo "Installing go-migrate"
	@cd src && go get
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


.PHONY: help
help:
	@printf "${COLOR_DIM} Available methods: ${COLOR_BASE} \n\n"
	@cat $(MAKEFILE_LIST) | \
	 	grep -E '^[a-zA-Z_]+:.* ## .*$$' | \
		sed -rn 's/(.*): ## (.*)/${COLOR_CMD}\1:${COLOR_DESCR}\2/p' | \
		column -t -s ":"
