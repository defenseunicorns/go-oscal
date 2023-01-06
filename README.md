# go-oscal

Repository for the generation of OSCAL data types

## Usage

Clone the repository and change into the `go-oscal` directory

Build the CLI

```bash
go build .
```

Generate Go structs from the OSCAL JSON schema:

```bash
./go-oscal --input test/oscal_component_schema.json \
           --output-file internal/oscal/types.go \
           --sub-struct \
           --pkg oscal
```

After running the above command, the auto-generated Go structs are output to a file at `internal/oscal/types.go`

To view the available CLI flags and their use-case:

```bash
./go-oscal --help
```

## Development

For development, the `Makefile` can be used to build, test, and generate the Go structs:

```bash
make
```
