# Notes

- Enhance logging and errors
- Abstract the logic to handle other OSCAL models/schemas
- Add/improve input validation for all command-line flags
  - --input-file: silent failures when a file path that does not exist is passed in
  - --pkg: should only accept valid syntax for go package names
  - --output-file: should we be creating non-existent directories if specified by user? or is documenting this sufficient?
  - --tags: should we be more flexible with how tags are passed in? currently only this format works - tag1,tag2 - or is documenting this sufficient?
