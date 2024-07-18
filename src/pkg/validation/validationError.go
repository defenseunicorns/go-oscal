package validation

import (
	"fmt"
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
func ExtractErrors(originalObject map[string]interface{}, detailedOutputUnit *jsonschema.OutputUnit) (validationErrors []ValidatorError) {
	if detailedOutputUnit == nil {
		return
	}

	topLevelErrors := make(map[string]*ValidatorError)
	flattenedErrors := Flatten(detailedOutputUnit)

	for _, schErr := range flattenedErrors {
		if schErr == nil || schErr.Error == nil {
			continue
		}
		var errorString string

		// Create the error string
		errorBytes, err := schErr.Error.MarshalJSON()
		if err != nil {
			errorString = fmt.Sprintf("error marshalling error: %v", err)
		} else {
			errorString = string(errorBytes)
		}

		// If the error is a top-level error, add it to the topLevelErrors map
		if schErr.InstanceLocation == "" {
			existing, ok := topLevelErrors[schErr.KeywordLocation]
			if !ok {
				topLevelErrors[schErr.KeywordLocation] = &ValidatorError{
					KeywordLocation:         schErr.KeywordLocation,
					AbsoluteKeywordLocation: schErr.AbsoluteKeywordLocation,
					InstanceLocation:        schErr.InstanceLocation,
					Error:                   errorString,
				}
			} else {
				existing.Error += ", " + errorString
			}
		} else {
			// Get the failed value
			failedValue := model.FindValue(originalObject, strings.Split(schErr.InstanceLocation, "/")[1:])
			// Skip if the failed value is a map or slice
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			// Create a ValidatorError from the jsonschema.BasicError
			validationError := ValidatorError{
				KeywordLocation:         schErr.KeywordLocation,
				AbsoluteKeywordLocation: schErr.AbsoluteKeywordLocation,
				InstanceLocation:        schErr.InstanceLocation,
				Error:                   errorString,
				FailedValue:             failedValue,
			}
			validationErrors = append(validationErrors, validationError)
		}
	}

	// Append the top level errors if there are no instance errors
	if len(validationErrors) == 0 {
		for _, value := range topLevelErrors {
			validationErrors = append(validationErrors, *value)
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
