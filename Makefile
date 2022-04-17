default: help

.PHONY: run
.PHONY: help
.PHONY: migrate
.PHONY: setup_dev


run: ## Run the server
	@cd src/ && go run cmd/server.go


migrate: ## Run migrations for both microservices
	@echo "Running migrations"
	@cd authentication && make migrate
	@cd user && make migrate

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
