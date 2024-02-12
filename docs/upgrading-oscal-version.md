# Upgrading Oscal Version
- When a new version is release an issue will be created on https://github.com/defenseunicorns/go-oscal/issues (see [tracking-oscal-version](./tracking-oscal-versions.md))
  
## Adding new Oscal version
- download the most recent undoctored `oscal_complete_schema.json` file at https://github.com/usnistgov/OSCAL/releases
- Run `make OSCAL_LATEST=X.X.X UNDOCTORED_SCHEMA=path/to/undoctored-schema.json upgrade`
    - This will run build, doctor the schema, placing it in `src/internal/schemas`, and generate the new types.
    - This will also delete the undoctored schema.