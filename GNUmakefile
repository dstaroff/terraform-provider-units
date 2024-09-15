default: help

##@ General

.PHONY: help
help: ## Show help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Code style

.PHONY: fmt
fmt: ## Format Go sources.
	gofmt -s -w -e -l .

.PHONY: lint
lint: ## Run linting against Go code.
	golangci-lint run

##@ Building

.PHONY: build
build: generate ## Build provider.
	go build -v ./...

.PHONY: install
install: build ## Install provider.
	go install -v ./...

##@ Testing

.PHONY: test
test: generate ## Run unit-tests.
	go test -v -cover -timeout=120s -parallel=6 ./...

.PHONY:test-acceptance
test-acceptance: generate ## Run acceptance tests.
	TF_ACC=1 go test -v -cover -timeout 120m ./...

##@ Misc

.PHONY: generate
generate: ## Generate source files.
	go generate ./internal/...

.PHONY: generate-docs
generate-docs: ## Format Terraform files in examples and generate docs.
	cd tools && go generate ./...
