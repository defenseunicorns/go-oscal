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

func GetOscalMapFromVersion(version string) (any, error) {
	schemaName := SCHEMA_PREFIX + strings.ReplaceAll(version, ".", "-")
	data, err := Schemas.ReadFile(schemaName + ".json")
	if err != nil {
		return nil, fmt.Errorf("failed to find schema: %w", err)
	}
	return jsonschema.UnmarshalJSON(bytes.NewReader(data))
}

func CreateSchemaPath(schemaVersion string) string {
	schemaName := SCHEMA_PATH + "/" + SCHEMA_PREFIX + strings.ReplaceAll(schemaVersion, ".", "-")
	return schemaName + ".json"
}

func GetOscalSchema(version string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()
	path := CreateSchemaPath(version)
	schemaMap, err := GetOscalMapFromVersion(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema map: %w", err)
	}
	compiler.AddResource(path, schemaMap)
	return compiler.Compile(path)
}
