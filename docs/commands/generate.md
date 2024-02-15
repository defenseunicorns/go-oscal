# [INTERNAL] Generate Oscal Types
The `generate` command is used internally to generate and export the individual types from a given `oscal_complete_schema.json`. It is designed specifically for use with the complete json schema's provided by [NIST OSCAL](https://github.com/usnistgov/OSCAL/releases) releases. In order to ensure consistency of types by oscal release, it is not recommended that this command be used except in the generation of newly released oscal version.

## Flags/Args
- `input-file` or `-f` is the filename of the OSCAL complete schema whose types should be generated.
- `output-file` or `-o` is the filename/path to write the generated oscal types to.
- `pkg` or `-p` is the package name for the generated types.
- `tags` or `-t` is the list of tags to be handled for marshalling and unmarshalling (default: json, yaml)