# go-oscal

Repository for the generation of OSCAL data types

## Usage

Clone the repository and change into the `go-oscal` directory

Build the CLI

```bash
go build .
```

To generate Go structs for the OSCAL Component Definition data model:

```bash
./go-oscal -f testdata/schema/component/oscal_component_schema.json \
           -o types.go
```

After running the above command, the auto-generated Go structs are output to a file in the root of the repository: `types.go`

To generate Go structs for the OSCAL System Security Plan data model:

```bash
./go-oscal -f testdata/schema/ssp/oscal_ssp_schema.json \
           -o types.go
```

After running the above command, the auto-generated Go structs are output to a file in the root of the repository: `types.go`

### Flags

The following command-line flags are available to use with `go-oscal`:

| Flag               | Description                                                                         |
|:-------------------|:------------------------------------------------------------------------------------|
| -h / --help        | lists all of the available flags and a short description of what they are used for  |
| -f / --input-file  | takes a path to an OSCAL JSON schema file                                           |
| -o / --output-file | the name of the file to save the Go structs to (outputs to STDOUT by default)       |
| -p / --pkg         | the name of the package for the generated code (defaults to "main")                 |
| -t / --tags        | comma seperated list of the tags to put on the structs (defaults to "json")         |

## Development

For development, the `Makefile` can be used to build, test, and generate the Go structs:

```bash
make
```
