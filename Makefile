.PHONY: all
all: clean build test generate-file

.PHONY: build
build: ## Build the binary
	go build .

.PHONY: clean
clean: ## Remove built binary
	go clean

.PHONY: generate-file
generate-file: ## Generate Go structs from OSCAL JSON schema and output to 'internal/oscal/types.go'
	./go-oscal --input-file test/oscal_component_schema.json --output-file internal/oscal/types.go --sub-struct --pkg oscal --tags json,yaml

.PHONY: generate-stdout
generate-stdout: ## Generate Go structs from OSCAL JSON schema and output to stdout
	./go-oscal --input-file test/oscal_component_schema.json --sub-struct --pkg oscal --tags json,yaml


.PHONY: test
test: ## Run unit tests
	go test -failfast -v ./...
