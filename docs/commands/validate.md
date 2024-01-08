# Validate existing OSCAL documents
The `validate` command can be used to validate an existing OSCAL document against the `oscal-version` declared in the documents `metadata` field. This is done by comparing the given document to the related `schema` field by field with the use of [github.com/santhosh-tekuri/jsonschema/v5](https://github.com/santhosh-tekuri/jsonschema). If the provided OSCAL document fails to `validate` against the `schema` a list of `errors` with the field location and reason for failure will be output to the provided `log-file` or `stderr` if no log file is provided.
* if a document fails validation, the command can be run with `-r` and a path to output the result that will provide an artifact with all fields that failed validation.

## Flags/Args
- `input-file` or `-f` is the filename containing the OSCAL model to be validated.
- `validation-result` or `-r` is the filename to write the `ValidationResult` `struct` to.
    - uses the file extension to determine the output format (`.json`/`.yaml`)
    - Provides information about the file that underwent validation and any errors that may have been encountered.

