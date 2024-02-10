package doctor_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/cmd/doctor"
	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
)

func TestDoctorCmd(t *testing.T) {
	t.Parallel()
	doctor.DoctorCmd.SetOut(new(bytes.Buffer))
	gooscaltest.RedirectLog(t)

	t.Run("returns an error if no inputfile is provided", func(t *testing.T) {
		doctor.DoctorCmd.SetArgs([]string{"-f", ""})
		err := doctor.DoctorCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if the inputfile is not a json file", func(t *testing.T) {
		doctor.DoctorCmd.SetArgs([]string{"-f", "../../../testdata/validation/assessment-result.yaml"})
		err := doctor.DoctorCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("writes the doctor output to the specified file", func(t *testing.T) {
		temp := t.TempDir()
		unDoctoredPath := "../../../testdata/doctor/oscal_complete_schema-"
		basePath := "../../../testdata/doctor/js-baseline/oscal_complete_schema-" // used to compare the output to the expected output
		outputPrefix := temp + "/doctored-"

		testCases := []string{"1-0-4.json", "1-1-1.json"}
		for _, testCase := range testCases {
			doctor.DoctorCmd.SetArgs([]string{"-f", unDoctoredPath + testCase, "-o", outputPrefix + testCase})

			// golang marshal uses recursion so the order of the keys is not guaranteed to match the expected output from the javascript implementation
			expected, err := os.ReadFile(basePath + testCase)
			if err != nil {
				t.Fatal(err)
			}
			// unmarshal and remarshal to get a consistent order of keys
			var expectedMap map[string]interface{}
			err = json.Unmarshal(expected, &expectedMap)
			if err != nil {
				t.Fatal(err)
			}
			expected, err = json.MarshalIndent(expectedMap, "", "  ")
			if err != nil {
				t.Fatal(err)
			}

			// run the doctor command
			err = doctor.DoctorCmd.Execute()
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}

			// read the output file and compare to the expected output
			actual, err := os.ReadFile(outputPrefix + testCase)
			if err != nil {
				t.Fatal(err)
			}

			if string(expected) != string(actual) {
				t.Errorf("Expected %s, got %s", expected, actual)
			}
		}

	})

}
