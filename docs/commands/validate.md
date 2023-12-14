# Validate existing OSCAL documents
The `validate` command can be used to validate an existing OSCAL document against the `oscal-version` declared in the documents `metadata` field. This is done by comparing the given document to the related `schema` field by field with the use of [github.com/santhosh-tekuri/jsonschema/v5](https://github.com/santhosh-tekuri/jsonschema). If the provided OSCAL document fails to `validate` against the `schema` a list of `errors` with the field location and reason for failure will be output to the provided `log-file` or `stderr` if no log file is provided.

## Flags/Args
- `input-file` or `-f` is the filename containing the OSCAL model to be validated.


