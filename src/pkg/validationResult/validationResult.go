package validationResult

import (
	"time"

	"github.com/defenseunicorns/go-oscal/src/pkg/validationError"
)

type ValidationResult struct {
	Success   bool                             `json:"success" yaml:"success"`
	TimeStamp time.Time                        `json:"timeStamp" yaml:"timeStamp"`
	Errors    []validationError.ValidatorError `json:"errors,omitempty" yaml:"errors,omitempty"`
}

func CreateValidationResult(errors []validationError.ValidatorError) ValidationResult {
	success := len(errors) == 0

	return ValidationResult{
		Success:   success,
		TimeStamp: time.Now(),
		Errors:    errors,
	}
}
