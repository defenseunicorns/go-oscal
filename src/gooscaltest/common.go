package gooscaltest

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"
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
