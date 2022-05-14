default: help

.PHONY: run
.PHONY: help
.PHONY: build
.PHONY: docker
.PHONY: migrate
.PHONY: create_migration
.PHONY: revert_migrations
.PHONY: setup_dev


# Migrations
N?=1

# Docker
TAG=latest
BRANCH_TAG=`git describe --abbrev=0 --tags | sort | head -n1`
ifeq ($(strip $(BRANCH_TAG)),)
TAG="${BRANCH_TAG}"
endif
IMAGE=ppcamp/go-microservice-authentication:${TAG}


run: ## Run the server
	@cd src/ && go run cmd/server.go


build: ## Build the server locally
	@cd src/ && go build -race cmd/server.go


lint: ## Run linters to this project. Remember to run `make setup_dev`
	@echo "Running linters"
	@cd src && golangci-lint run ./...


docker: ## Create docker image
	@echo "Building ${IMAGE}"
	@docker build --no-cache -f Dockerfile -t ${IMAGE} .


migrate: ## Run migrations created with `make create_migration`. Remember to `make setup_dev`
	@echo "Running migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose up


create_migration: ## Create a new migration, e.g `name=teste make create_migration`. Remember to `make setup_dev`
	@echo "Creating migration"
	@migrate create -ext sql -dir migrations -seq ${name}


revert_migrations: ## Revert a given migration, e.g `N=2 make revert_migrations`, by default 1. Remember to `make setup_dev`
	@echo "Reverting migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose down ${N}



setup_dev: ## Install dev dependencies
	@echo "Installing go-migrate"
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "Installing linters"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2


help:
	@printf "\e[2m Available methods:\033[0m\n\n"
# 1. read makefile
# 2. get lines that can have a method description and assign colors to method
# 3. colour special worlds. If fail, return the original row
# 4. colour and strip lines
# 5. create column view
	@cat $(MAKEFILE_LIST) | \
	 	grep -E '^[a-zA-Z_]+:.* ## .*$$' | \
		sed -rn 's/`([a-zA-Z0-9=\_\ \-]+)`/\x1b[33m\1\x1b[0m/g;t1;b2;:1;h;:2;p' | \
		sed -rn 's/(.*): ## (.*)/\x1b[32m\1:\x1b[0m\2/p' | \
		column -t -s ":"
