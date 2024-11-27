# Generating OSCAL Types

This project uses a custom script to generate Go types from OSCAL JSON schemas. The process is automated using the `go generate` command.

## Prerequisites

- Bash shell
- Node.js and npm (for running `npx quicktype`)
- Go installed on your system

## Generation Process

1. The type generation is triggered by running:

   ```
   go generate
   ```

   This command is defined in the `main.go` file:

   ```go
   //go:generate ./hack/gen-types.sh
   ```

2. The `gen-types.sh` script performs the following actions:
- Iterates through all `oscal_complete_schema-*.json` files in the `src/internal/schemas` directory
- Extracts the version number from each filename
- Creates an output directory for each version
- Uses `quicktype` to generate Go types from the JSON schema
- Adds YAML and XML tags to the generated structs using `sed`

## Output

The generated types are placed in `src/types/oscal-<version>/types.go`, where `<version>` is the OSCAL schema version (e.g., `1-0-0` for version 1.0.0).

Each generated file contains a package named `oscalTypes_<version>` and a top-level struct named `OscalModels`.

## Updating Types

To update the generated types:

1. Ensure you have the latest OSCAL JSON schemas in `src/internal/schemas`
2. Run `go generate` in the project root directory

This will regenerate all type files based on the current schemas.

