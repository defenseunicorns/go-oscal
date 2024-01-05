package validation

import (
	"time"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
)

type ValidationResult struct {
	Valid     bool                     `json:"valid" yaml:"valid"`
	TimeStamp time.Time                `json:"timeStamp" yaml:"timeStamp"`
	Errors    []ValidatorError         `json:"errors,omitempty" yaml:"errors,omitempty"`
	Metadata  ValidationResultMetadata `json:"metadata" yaml:"metadata"`
}

type ValidationResultMetadata struct {
	DocumentType    string `json:"documentType" yaml:"documentType"`
	DocumentVersion string `json:"documentVersion" yaml:"documentVersion"`
	SchemaVersion   string `json:"schemaVersion" yaml:"schemaVersion"`
}

// NewValidationResult creates a new ValidationResult from a Validator and a slice of ValidatorErrors
func NewValidationResult(validator *Validator, errors []ValidatorError) ValidationResult {
	documentVersion, err := utils.GetOscalVersionFromMap(validator.GetJsonModel())
	if err != nil {
		documentVersion = ""
	}
	return ValidationResult{
		Valid:     len(errors) == 0,
		TimeStamp: time.Now(),
		Errors:    errors,
		Metadata: ValidationResultMetadata{
			DocumentType:    validator.GetModelType(),
			DocumentVersion: documentVersion,
			SchemaVersion:   validator.GetSchemaVersion(),
		},
	}
}
