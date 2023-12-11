package generate

import (
	"bytes"
	"strings"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
)

func TestGenerateCmd(t *testing.T) {
	t.Parallel()
	testInput := "../../../schema/complete/oscal_complete_schema-1-0-4.json"
	packageName := "oscalTypes"
	tags := "json,yaml"
	GenerateCmd.SetOut(new(bytes.Buffer))

	logOutput := gooscaltest.RedirectLog(t)

	t.Run("baseline", func(t *testing.T) {
		expectedValues := []string{"Catalog", "Profile", "ComponentDefinition", "SystemSecurityPlan", "AssessmentPlan", "AssessmentResults", "PlanOfActionAndMilestones", "yaml:", "json"}

		GenerateCmd.SetArgs([]string{"-f", testInput, "-t", tags, "-p", packageName})
		err := GenerateCmd.Execute()
		if err != nil {
			t.Error("expected nil, got", err)
		}
		actual := string(gooscaltest.ReadLog(t, logOutput))

		t.Run("generates a go types document that gets logged if no output file is provided", func(t *testing.T) {
			for _, val := range expectedValues {
				if !strings.Contains(actual, val) {
					t.Error("expected", val, "got", actual)
				}
			}
		})

		t.Run("sets the package name using the -p flag", func(t *testing.T) {
			if !strings.Contains(actual, "package "+packageName) {
				t.Error("expected package oscalTypes, got", actual)
			}
		})
	})

	t.Run("input-file", func(t *testing.T) {

		t.Run("returns an error if no inputfile is provided", func(t *testing.T) {
			GenerateCmd.SetArgs([]string{"-f", ""})
			err := GenerateCmd.Execute()
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
		t.Run("returns an error if it fails to read the input file", func(t *testing.T) {
			GenerateCmd.SetArgs([]string{"-f", "non-existant.json"})
			err := GenerateCmd.Execute()
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	})

	t.Run("output-file", func(t *testing.T) {
		t.Run("outputs to a file if the -o flag is provided", func(t *testing.T) {
			outputFile := "test_output.go"
			GenerateCmd.SetArgs([]string{"-f", testInput, "-o", outputFile})
			err := GenerateCmd.Execute()
			if err != nil {
				t.Error("expected nil, got", err)
			}
		})
	})
}
