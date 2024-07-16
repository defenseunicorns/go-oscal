package validation

import (
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Extension of the jsonschema.BasicError struct to include the failed value
// if the failed value is a map or slice, it will be omitted
type ValidatorError struct {
	// KeywordLocation is the location of the keyword in the schema for failing value
	KeywordLocation string `json:"keywordLocation" yaml:"keywordLocation"`
	// AbsoluteKeywordLocation is the absolute location of the keyword in the schema for failing value
	AbsoluteKeywordLocation string `json:"absoluteKeywordLocation" yaml:"absoluteKeywordLocation"`
	// InstanceLocation is the location of the instance in the document
	InstanceLocation string `json:"instanceLocation" yaml:"instanceLocation"`
	// Error is the error message
	Error string `json:"error" yaml:"error"`
	// FailedValue is the value of the key that failed validation
	FailedValue interface{} `json:"failedValue,omitempty" yaml:"failedValue,omitempty"`
}

// Creates a []ValidatorError from a jsonschema.Basic
// The jsonschema.Basic contains the errors from the validation
func ExtractErrors(originalObject map[string]interface{}, validationError jsonschema.ValidationError) (validationErrors []ValidatorError) {
	for _, cause := range validationError.Causes {
		basicOutput := cause.BasicOutput()
		basicError := cause.Error()

		if !strings.HasPrefix(basicError, "missing properties:") && (basicError == "" || strings.HasPrefix(basicError, "doesn't validate with")) {
			continue
		}
		if len(validationErrors) > 0 && validationErrors[len(validationErrors)-1].InstanceLocation == basicOutput.InstanceLocation {
			validationErrors[len(validationErrors)-1].Error += ", " + basicError
		} else {
			failedValue := model.FindValue(originalObject, strings.Split(basicOutput.InstanceLocation, "/")[1:])
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			// Create a ValidatorError from the jsonschema.BasicError
			validationError := ValidatorError{
				KeywordLocation:         basicOutput.KeywordLocation,
				AbsoluteKeywordLocation: basicOutput.AbsoluteKeywordLocation,
				InstanceLocation:        basicOutput.InstanceLocation,
				Error:                   basicError,
				FailedValue:             failedValue,
			}
			validationErrors = append(validationErrors, validationError)
		}
	}
	return validationErrors

}
