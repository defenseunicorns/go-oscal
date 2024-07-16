package schemas

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed *.json
var Schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
)

type TempFile struct {
	Path    string
	Cleanup func()
}

// CreateTempFile creates a temporary file for the given schema and returns the file path and a cleanup function.
func CreateTempFile(schemaVersion string) (*TempFile, error) {
	schemaName := SCHEMA_PREFIX + strings.ReplaceAll(schemaVersion, ".", "-")
	schemaPath := schemaName + ".json"
	data, err := Schemas.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to find schema: %w", err)
	}

	tempFile, err := os.CreateTemp("", schemaName+"-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := tempFile.Write(data); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("failed to write to temp file: %w", err)
	}

	cleanup := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return &TempFile{
		Path:    tempFile.Name(),
		Cleanup: cleanup,
	}, nil
}

type SchemaLoader struct {
	fs embed.FS
}

// NewSchemaLoader creates a new schema loader from the default schema file.
func NewSchemaLoader() *SchemaLoader {
	return SchemaLoaderFromEmbedFS(Schemas)
}

// SchemaLoaderFromEmbedFS creates a new schema loader from an embed.FS.
func SchemaLoaderFromEmbedFS(fs embed.FS) *SchemaLoader {
	return &SchemaLoader{
		fs: fs,
	}
}

func (s *SchemaLoader) Load(url string) (any, error) {
	urlParts := strings.Split(url, "/")
	schemaVersion := urlParts[len(urlParts)-1]
	schemaName := SCHEMA_PREFIX + strings.ReplaceAll(schemaVersion, ".", "-")
	schemaPath := schemaName + ".json"
	data, err := s.fs.Open(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to find schema: %w", err)
	}
	return jsonschema.UnmarshalJSON(data)
}
