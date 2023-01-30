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
./go-oscal --input-file test/oscal_component_schema.json \
           --output-file types.go \
           --tags json,yaml
```

After running the above command, the auto-generated Go structs are output to a file in the root of the repository: `types.go`

To view the available CLI flags and their use-case:

```bash
./go-oscal --help
```

## Development

***Note***: The unit tests are not in a passing state currently. There is an [issue](https://github.com/defenseunicorns/go-oscal/issues/9) to address this

For development, the `Makefile` can be used to build, test, and generate the Go structs:

```bash
make
```
