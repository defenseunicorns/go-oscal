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
	SCHEMA_PATH   = "https://raw.githubusercontent.com/defenseunicorns/go-oscal/main/src/internal/schemas/"
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

// GetOscalSchema returns the jsonschema Schema for a given version
func GetOscalSchema(version string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()

	// Create the schemaUrl (url) to the github schema file
	// URL is used in ValidationError to point to the exact schema that failed validation
	schemaUrl := CreateSchemaPath(version)
	// Get the schema map from the local schema file
	schemaMap, err := GetOscalMapFromVersion(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema map: %w", err)
	}

	// Add the schema map to the compiler with the url
	compiler.AddResource(schemaUrl, schemaMap)
	// Compile the schema
	return compiler.Compile(schemaUrl)
}
