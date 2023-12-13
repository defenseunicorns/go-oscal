# Convert existing oscal documents to another version
The `convert` command is used to upgrade an existing oscal document to another version. If it fails to upgrade, the errors that require fixing will be outputted to the provided log file (if provided) or `STDERR`

## Flags/Args

- `input-file` or `-f` is the filename containing the OSCAL model to be upgraded.
    - Supports `.json` or `.yaml`
- `output-file` or `-o` is the filename to write the upgraded OSCAL model to.
    - if successful the upgraded model will be written to this file
    - uses the file extension to determine the output format (`.json`/`.yaml`)
- `version` or `-v` the version to be upgraded to.
    - will log error if the version is not supported or is not in the valid `1.0.4` format.