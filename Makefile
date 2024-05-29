# PREAMBLE
#//////////////////////////////////////////////////////////////////////////////
#
SHELL := bash

# VARIABLES, CONFIG, & SETTINGS
#//////////////////////////////////////////////////////////////////////////////
#
BINDIR       := $(CURDIR)/bin
BINNAME      ?= go-oscal
INSTALL_PATH ?= /usr/local/bin
# This should be replaced with the latest version of the OSCAL schema.
OSCAL_LATEST ?= 1-1-2
# This should be replaced with the path to the latest oscal complete json schema associated with OSCAL_LATEST.
UNDOCTORED_SCHEMA ?= testdata/doctor/oscal_complete_schema-1-1-1.json
OSCAL_LATEST_SCHEMA := src/internal/schemas/oscal_complete_schema-$(OSCAL_LATEST).json
OSCAL_LATEST_OUTPUT := src/types/oscal-$(OSCAL_LATEST)/
OSCAL_LATEST_PACKAGE := oscalTypes_$(subst -,_,$(OSCAL_LATEST))

# Git vars
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)


# Go CLI options
PKG         := ./...
TAGS        :=
TESTS       := .
TESTFLAGS   := -race -v -failfast
LDFLAGS     := -w -s
GOFLAGS     :=
CGO_ENABLED ?= 0

# Allows us to set VERSION from the command line.
# Otherwise, if BINARY_VERSION is not set, use the current git tag.
ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

# Project sources.
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

# TASKS
#//////////////////////////////////////////////////////////////////////////////
#
.PHONY: all
all: clean build test generate-compdef-stdout generate-ssp-stdout

.PHONY: help
help: ## Show this help message.
	@grep -E '^[a-zA-Z_/-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "; printf "\nUsage:\n"}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo

.PHONY: clean
clean: ## Remove generated artifacts.
	go clean
	rm -rf $(BINDIR)

.PHONY: build
build: $(BINDIR)/$(BINNAME) ## Build the project.

$(BINDIR)/$(BINNAME): $(SRC)
	CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)/$(BINNAME)' .

.PHONY: test
test: build ## Run automated tests.
	go clean -testcache && go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: install
install: ## Install binary to $INSTALL_PATH.
	@install "$(BINDIR)/$(BINNAME)" "$(INSTALL_PATH)/$(BINNAME)"

.PHONY: upgrade
upgrade: ## Upgrade oscal schema version and generate new types and doctored schema. 
	make doctor-latest-schema
	make generate-latest-schema
	echo -e "---\noscal: v$(subst -,.,$(OSCAL_LATEST))" > update/oscal-version.yaml
	rm $(UNDOCTORED_SCHEMA)

.PHONY: doctor-latest
doctor-latest-schema: clean build
	$(BINDIR)/$(BINNAME) doctor -f $(UNDOCTORED_SCHEMA) -o $(OSCAL_LATEST_SCHEMA)

.PHONY: generate-latest
generate-latest-schema: clean build
	$(BINDIR)/$(BINNAME) generate -f $(OSCAL_LATEST_SCHEMA) --pkg $(OSCAL_LATEST_PACKAGE) --tags json,yaml -o $(OSCAL_LATEST_OUTPUT)/types.go