default: help

.PHONY: run
.PHONY: help
.PHONY: migrate
.PHONY: generate
.PHONY: setup_dev
.PHONY: proto
.PHONY: certificate

JS_IMPORT_STYLE = import_style=commonjs,binary
JS_OUT_STYLE = import_style=typescript,mode=grpcwebtext
JS_OUT_DIR = generated/web
PROTO_DIR = protos


run: ## Run the server
	@cd src/ && go run cmd/server.go


migrate: ## Run migrations for both microservices
	@echo "Running migrations"
	@cd authentication && make migrate
#   @cd user && make migrate


setup_dev: ## Install dev dependencies
	@echo "Installing protobuf"
	@sudo apt install protobuf-compiler
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


proto: ## Generate grpc-web protos
	@echo "Generating grpc-web protos"
	@mkdir -p $(JS_OUT_DIR)
	@protoc -I=$(PROTO_DIR) $(PROTO_DIR)/* \
    	--js_out=$(JS_IMPORT_STYLE):$(JS_OUT_DIR) \
    	--grpc-web_out=$(JS_OUT_STYLE):$(JS_OUT_DIR)


dcup: ## Start docker compose
	@docker-compose up --force-recreate --build -d


dcdown: ## Close docker compose instances
	@docker-compose down


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
		sed -rn 's/(.*):.* ## (.*)/\x1b[32m\1:\x1b[0m\2/p' | \
		column -t -s ":"
