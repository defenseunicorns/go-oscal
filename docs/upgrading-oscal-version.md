# Upgrading an OSCAL Version

When a new OSCAL version is released, Renovate opens an update PR and the OSCAL-version workflow creates an issue. See [tracking OSCAL versions](./tracking-oscal-versions.md).

## Prerequisites

1. Confirm the Renovate PR updates both version declarations:
   - `update/oscal-version.yaml` (`oscal: vX.Y.Z`)
   - `src/pkg/versioning/versioning.go` (`latestVersion = "X.Y.Z"`)

   `TestGetLatestVersion` requires these values to match.
2. Update `OSCAL_LATEST` in `Makefile` to `X-Y-Z`.
3. Download the release's unmodified `oscal_complete_schema.json` asset from the [NIST OSCAL releases](https://github.com/usnistgov/OSCAL/releases) page to a disposable path. Record the release URL and SHA-256 checksum in the PR for provenance.

## Generate the supported schema and types

Run:

```sh
make OSCAL_LATEST=X-Y-Z \
  UNDOCTORED_SCHEMA=/tmp/oscal_complete_schema-X-Y-Z.json \
  upgrade
```

The target builds `go-oscal`, doctors the schema into `src/internal/schemas`, generates `src/types/oscal-X-Y-Z/types.go`, updates `update/oscal-version.yaml`, and deletes `UNDOCTORED_SCHEMA`. Do not pass the only retained copy of the release asset.

## Review the schema delta

Compare the new and previous NIST release assets and record:

- added or removed definitions and properties;
- changes to requiredness, cardinality, property type, references, and enum values;
- documentation-only changes; and
- whether each structural change is represented in the generated Go types.

Distinguish validation-only changes, such as an enum expansion on a generated `string` field, from changes that alter the generated Go struct surface.

## Verify the upgrade

1. Regenerate the added types and compare the result with the committed file:

   ```sh
   bin/go-oscal generate \
     --input-file src/internal/schemas/oscal_complete_schema-X-Y-Z.json \
     --pkg oscalTypes_X_Y_Z \
     --tags json,yaml \
     --output-file /tmp/oscal-X-Y-Z-types.go
   cmp /tmp/oscal-X-Y-Z-types.go src/types/oscal-X-Y-Z/types.go
   ```

2. Run the test target with CGO enabled because it uses the Go race detector:

   ```sh
   CGO_ENABLED=1 make test
   ```