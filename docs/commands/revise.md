# Revise existing oscal documents to another version
The `revise` command is used to validate and export a given OSCAL document as another version. This is currently done by taking in an existing OSCAL document, via the `input-file` flag, and running validation against the `schema` of the `oscal-version` requested. If the OSCAL document validates successfully against the requested `oscal-version` it is exported to the provided `output-file` or `stdout` if no file is provided. If the OSCAL document fails to validate against the requested `oscal-version` `schema`, the errors that need to be fixed in order to `revise` are logged to the provided `log-file` or `stderr` if none is provided.

## Flags/Args

- `input-file` or `-f` is the filename containing the OSCAL model to be upgraded.
    - Supports `.json` or `.yaml`
- `output-file` or `-o` is the filename to write the upgraded OSCAL model to.
    - if successful the upgraded model will be written to this file
    - uses the file extension to determine the output format (`.json`/`.yaml`)
- `version` or `-v` the version to be upgraded to.
    - will log error if the version is not supported or is not in the valid format (ie `1.0.4`).