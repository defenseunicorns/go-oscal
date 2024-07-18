package schemas

import (
	"bytes"
	"embed"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed *.json
var Schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
	SCHEMA_PATH   = "https://github.com/defenseunicorns/go-oscal/tree/main/src/internal/schemas"
)

// GetOscalMapFromVersion returns the OSCAL schema map for a given version
func GetOscalMapFromVersion(version string) (any, error) {
	schemaName := SCHEMA_PREFIX + strings.ReplaceAll(version, ".", "-")
	data, err := Schemas.ReadFile(schemaName + ".json")
	if err != nil {
		return nil, fmt.Errorf("failed to find schema: %w", err)
	}
	return jsonschema.UnmarshalJSON(bytes.NewReader(data))
}

// CreateSchemaPath creates a path to the schema file for a given version
func CreateSchemaPath(schemaVersion string) string {
	schemaName := SCHEMA_PATH + "/" + SCHEMA_PREFIX + strings.ReplaceAll(schemaVersion, ".", "-")
	return schemaName + ".json"
}
