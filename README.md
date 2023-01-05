# go-oscal

Repository for the generation of OSCAL data types 

## Usage

Clone the repository and  change into the `go-oscal` directory

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

To view the available CLI flags and their use-case:

```bash
./go-oscal --help
```
