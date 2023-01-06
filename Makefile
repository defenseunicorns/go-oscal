.PHONY: all
all: clean build test generate

.PHONY: build
build: ## Build the binary
	go build .

.PHONY: clean
clean: ## Remove built binary
	go clean

.PHONY: generate
generate: ## Generate Go structs from OSCAL JSON schema and output to 'internal/oscal/types.go'
	./go-oscal --input test/oscal_component_schema.json \
	           --output-file internal/oscal/types.go \
	           --sub-struct \
	           --pkg oscal

.PHONY: test
test: ## Run unit tests
	go test -v ./internal/oscal
