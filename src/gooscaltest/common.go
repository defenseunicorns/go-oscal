package gooscaltest

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	byteMapMtx                      = sync.Mutex{}
	logMtx                          = sync.Mutex{}
	ByteMap                         = map[string][]byte{}
	ValidComponentPath              = "../../../testdata/validation/valid-component-definition.yaml"
	NoVersionComponentPath          = "../../../testdata/validation/no-version-component-definition.yaml"
	InvalidVersionComponentPath     = "../../../testdata/validation/invalid-version-component-definition.yaml"
	UnsupportedVersionComponentPath = "../../../testdata/validation/unsupported-version-component-definition.yaml"
	ValidAssessmentResultPath       = "../../../testdata/validation/assessment-result.yaml"
	ValidCatalogPath                = "../../../testdata/validation/catalog.yaml"
	ValidCatalogJsonPath            = "../../../testdata/validation/catalog.json"
	InvalidCatalogPath              = "../../../testdata/validation/invalid-catalog.yaml"
	ValidProfilePath                = "../../../testdata/validation/profile.yaml"
	ValidSSP                        = "../../../testdata/validation/system-security-plan.yaml"
	ValidPlanOfActionAndMilestones  = "../../../testdata/validation/plan-of-action-and-milestones.json"
	ValidAsessmentPlan              = "../../../testdata/validation/assessment-plan.json"
	MultipleDocPath                 = "../../../testdata/validation/multiple.yaml"
	writers                         = []io.Writer{}

	pathSlice = []string{ValidComponentPath, NoVersionComponentPath, InvalidVersionComponentPath, UnsupportedVersionComponentPath, ValidAssessmentResultPath, ValidCatalogPath, ValidCatalogJsonPath, InvalidCatalogPath, ValidProfilePath, ValidSSP, MultipleDocPath, ValidPlanOfActionAndMilestones, ValidAsessmentPlan}
)

// GetOscalSchema returns the jsonschema Schema for a given version
func GetOscalSchema(version string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()

	// Create the schemaUrl (url) to the github schema file
	// URL is used in ValidationError to point to the exact schema that failed validation
	schemaUrl := schemas.CreateSchemaPath(version)
	// Get the schema map from the local schema file
	schemaMap, err := schemas.GetOscalMapFromVersion(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema map: %w", err)
	}

	// Add the schema map to the compiler with the url
	compiler.AddResource(schemaUrl, schemaMap)
	// Compile the schema
	return compiler.Compile(schemaUrl)
}

// GetByteMap reads the files in PathSlice and stores them in ByteMap
func GetByteMap(t *testing.T) {
	byteMapMtx.Lock()
	defer byteMapMtx.Unlock()
	if len(ByteMap) == 0 {
		for _, path := range pathSlice {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			ByteMap[path] = bytes
		}
	}
}

func RedirectLog(t *testing.T) *bytes.Buffer {
	logMtx.Lock()
	defer logMtx.Unlock()
	logOutput := new(bytes.Buffer)
	writers = append(writers, logOutput)
	multiWriter := io.MultiWriter(writers...)
	log.SetOutput(multiWriter)

	return logOutput
}

func ReadLog(t *testing.T, logOutput *bytes.Buffer) []byte {
	logMtx.Lock()
	defer logMtx.Unlock()
	return logOutput.Bytes()
}
