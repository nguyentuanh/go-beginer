IGNORE_DIR1=-path ./datastorage -prune -o
GOFMT_FILES?=$$(find . $(IGNORE_DIR1) -name '*.go' | grep -v vendor)
GOFMT := "goimports"

fmt: ## Run gofmt for all .go files
	@$(GOFMT) -w -d -e -l $(GOFMT_FILES)

DEPEND=\
	golang.org/x/tools/cmd/goimports \
	golang.org/x/tools/cmd/stringer \
	github.com/swaggo/swag/cmd/swag \
    github.com/githubnemo/CompileDaemon

depend: ## Install dependencies for dev
	@go get -v ./...
	@go get -v $(DEPEND)

lint: ## run linter
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.33.0 golangci-lint run -v

dev:
	scripts/run_local.sh

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
