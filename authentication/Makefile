default: help

include .env

# COLORS
COLOR_METHOD=\033[36m
COLOR_DEFAULT=\033[0m
COLOR_DIM=\e[2m
N?=1

.PHONY: run_server
run_server: ## Run the server
	@cd src/ && go run cmd/*.go


.PHONY: generate
generate: ## Create golang protos
	@echo "Generating protos"
	@echo " - Generating messages"
	@protoc --go_out=. protos/*.proto
	@echo " - Generating services"
	@protoc --go-grpc_out=. protos/*.proto


.PHONY: migrate
migrate: ## Run migrations
	@echo "Running migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose up

.PHONY: revert_migrations
revert_migrations: ## Revert a given migration, e.g N=2 make revert_migrations, by default 1
	@echo "Reverting migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose down ${N}


.PHONY: setup_dev
setup_dev: ## Install dev dependencies
	@echo "Installing go-migrate"
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


.PHONY: create_migration
create_migration: ## Create a new migration, e.g name=teste make create_migration
	@echo "Creating migration"
	@migrate create -ext sql -dir migrations -seq ${name}


.PHONY: help
help:
	@printf "${COLOR_DIM} Available methods: ${COLOR_DEFAULT} \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort  | \
		awk \
			'BEGIN {FS = ":.*?## "}; { \
				printf "${COLOR_METHOD}%s: ${COLOR_DEFAULT}%s\n", $$1, $$2 \
			}'
