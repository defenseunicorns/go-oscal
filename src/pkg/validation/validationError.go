package validation

import (
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Extension of the jsonschema.BasicError struct to include the failed value
// if the failed value is a map or slice, it will be omitted
type ValidatorError struct {
	jsonschema.BasicError
	FailedValue interface{} `json:"failedValue,omitempty" yaml:"failedValue,omitempty"`
}

// Creates a []ValidatorError from a jsonschema.Basic
// The jsonschema.Basic contains the errors from the validation
func ExtractErrors(originalObject map[string]interface{}, validationError jsonschema.Basic) (validationErrors []ValidatorError) {
	for _, error := range validationError.Errors {
		if error.InstanceLocation == "" || error.Error == "" || strings.HasPrefix(error.Error, "doesn't validate with") {
			continue
		}
		if len(validationErrors) > 0 && validationErrors[len(validationErrors)-1].InstanceLocation == error.InstanceLocation {
			validationErrors[len(validationErrors)-1].Error += ", " + error.Error
		} else {
			failedValue := utils.FindValue(originalObject, strings.Split(error.InstanceLocation, "/")[1:])
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			validationErrors = append(validationErrors, ValidatorError{BasicError: error, FailedValue: failedValue})
		}
	}
	return validationErrors

}
