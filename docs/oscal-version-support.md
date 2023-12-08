# How to add support for a new version of OSCAL

As noted in how we [track OSCAL updates](./tracking-oscal-versions.md), an issue will be created by renovate/GitHub actions when a new version of OSCAL is released.

## Steps

1. Create a new branch for the new version of OSCAL.
2. Download the `oscal_complete_schema.json` file from the release.
3. Ensure the JSON is valid. Run the schema through the [OSCAL JSON Doctor](https://github.com/defenseunicorns/oscal-json-doctor).
4. Add the new schema json file to `src/pkg/validation/schema` with a version identifier that matches the version of OSCAL.
5. Add the version to the `supportedVersion` variable in `src/pkg/validation/utils.go`.
6. Create a new directory `src/types/oscal-x-x-x/`
7. Build and generate the new types with `make build && ./bin/go-oscal generate -f src/pkg/validation/schema/oscal_complete_schema-x-x-x.json -o src/types/oscal-x-x-x/types.go -p oscalTypes -t json,yaml`.
8. Add, commit, and push the changes to the new branch and submit a Pull Request.

## Testing

TODO: when logic is present to upgrade an older version document to a newer version - I'd like to add tests that always run against the latest version.

## Notes
When we have the upgrade logic - we'll likely need to add more steps to this process to account for the delta between versions and any programmatic adaptation required. Generically we will want to evolve this and the testing between different logic to be aligned in-support-of new versions of OSCAL. 