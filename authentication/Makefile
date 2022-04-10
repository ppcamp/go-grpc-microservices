default: help

# COLORS
COLOR_METHOD=\033[36m
COLOR_DEFAULT=\033[0m
COLOR_DIM=\e[2m


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


.PHONY: help
help:
	@printf "${COLOR_DIM} Available methods: ${COLOR_DEFAULT} \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort  | \
		awk \
			'BEGIN {FS = ":.*?## "}; { \
				printf "${COLOR_METHOD}%s\t\t${COLOR_DEFAULT}%s\n", $$1, $$2 \
			}'
