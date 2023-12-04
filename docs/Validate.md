# Validate existing OSCAL documents
The `validate` command can be used to validate an existing OSCAL document against the `oscal-version` declared in the documents `metadata` field.

## Flags/Args
- `input-file` or `-f` is the filename containing the OSCAL model to be validated.
- `log-file` or `l` is the name of the file to output to
    - optional: defaults to `STDERR`
    - Log open-time flags: `os.O_CREATE|os.O_WRONLY|os.O_APPEND`
    - Schema validation failures are logged in `JSON` format. 

