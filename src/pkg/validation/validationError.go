package validation

import (
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/santhosh-tekuri/jsonschema/v5"
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
func ExtractErrors(originalObject map[string]interface{}, validationError jsonschema.Basic) (validationErrors []ValidatorError) {
	for _, basicError := range validationError.Errors {
		if !strings.HasPrefix(basicError.Error, "missing properties:") && (basicError.InstanceLocation == "" || basicError.Error == "" || strings.HasPrefix(basicError.Error, "doesn't validate with")) {
			continue
		}
		if len(validationErrors) > 0 && validationErrors[len(validationErrors)-1].InstanceLocation == basicError.InstanceLocation {
			validationErrors[len(validationErrors)-1].Error += ", " + basicError.Error
		} else {
			failedValue := model.FindValue(originalObject, strings.Split(basicError.InstanceLocation, "/")[1:])
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			// Create a ValidatorError from the jsonschema.BasicError
			validationError := ValidatorError{
				KeywordLocation:         basicError.KeywordLocation,
				AbsoluteKeywordLocation: basicError.AbsoluteKeywordLocation,
				InstanceLocation:        basicError.InstanceLocation,
				Error:                   basicError.Error,
				FailedValue:             failedValue,
			}
			validationErrors = append(validationErrors, validationError)
		}
	}
	return validationErrors

}
