package gooscaltest

import (
	"os"
	"sync"
	"testing"
)

var (
	mutex                           = sync.Mutex{}
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
	BasicErrorPathJson              = "../../../testdata/validation/basic-error.json"

	pathSlice = []string{ValidComponentPath, NoVersionComponentPath, InvalidVersionComponentPath, UnsupportedVersionComponentPath, ValidAssessmentResultPath, ValidCatalogPath, ValidCatalogJsonPath, InvalidCatalogPath, ValidProfilePath, ValidSSP, MultipleDocPath, ValidPlanOfActionAndMilestones, ValidAsessmentPlan}
)

// GetByteMap reads the files in PathSlice and stores them in ByteMap
func GetByteMap(t *testing.T) {
	mutex.Lock()
	if len(ByteMap) == 0 {
		for _, path := range pathSlice {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			ByteMap[path] = bytes
		}
	}
	mutex.Unlock()
}
