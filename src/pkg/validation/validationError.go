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
func ExtractErrors(originalObject map[string]interface{}, validationError jsonschema.OutputUnit) (validationErrors []ValidatorError) {
	// Flatten the output unit to get all the errors
	flattenedErrors := Flatten(&validationError)

	for _, cause := range flattenedErrors {
		// Skip nil errors and empty errors (nested)
		if cause == nil || cause.Error == nil {
			continue
		}

		// Create the error string
		errorBytes, err := cause.Error.MarshalJSON()
		if err != nil {
			continue
		}
		basicError := string(errorBytes)

		// Skip empty errors and errors starting with "missing properties:" and "doesn't validate with"
		if !strings.HasPrefix(basicError, "missing properties:") && (basicError == "" || strings.HasPrefix(basicError, "doesn't validate with")) {
			continue
		}

		// Append the error to the last error if it has the same instance location
		if len(validationErrors) > 0 && validationErrors[len(validationErrors)-1].InstanceLocation == cause.InstanceLocation {
			validationErrors[len(validationErrors)-1].Error += ", " + basicError
		} else {
			// Get the failed value
			failedValue := model.FindValue(originalObject, strings.Split(cause.InstanceLocation, "/")[:1])
			// Skip if the failed value is a map or slice
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			// Create a ValidatorError from the jsonschema.BasicError
			validationError := ValidatorError{
				KeywordLocation:         cause.KeywordLocation,
				AbsoluteKeywordLocation: cause.AbsoluteKeywordLocation,
				InstanceLocation:        cause.InstanceLocation,
				Error:                   basicError,
				FailedValue:             failedValue,
			}
			validationErrors = append(validationErrors, validationError)
		}
	}
	return validationErrors
}

// Flatten function to collect all OutputUnits into a slice
func Flatten(outputUnit *jsonschema.OutputUnit) []*jsonschema.OutputUnit {
	var result []*jsonschema.OutputUnit
	var flattenHelper func(*jsonschema.OutputUnit)

	flattenHelper = func(unit *jsonschema.OutputUnit) {
		if unit == nil {
			return
		}
		result = append(result, unit)
		for _, child := range unit.Errors {
			flattenHelper(&child)
		}
	}

	flattenHelper(outputUnit)
	return result
}
