# [INTERNAL] Doctor json schemas
The `doctor` command is used internally to fix known issues with an `oscal_complete_schema.json` so that it may be properly compiled for use in the `validate` and `revise` commands. This command should be used after the release of a new [OSCAL VERSION](https://github.com/usnistgov/OSCAL/releases) in order to doctor and output a new schema into `src/internal/schemas/`.

## Current Fixes
- `$id` and `$ref` being defined in the same definition.
    - This is handled by following the ref and "lifting" the definition up to referrer. (jsonschema v7 does not support $id and $ref both being defined)
  
## Flags/Args
- `input-file` or `-f` is the path to the OSCAL complete json schema to be doctored.
- `output-file` or `-o` is the path to write the doctored schema to. 