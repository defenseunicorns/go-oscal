package revision

import (
	"encoding/json"
	"fmt"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"gopkg.in/yaml.v3"
)

type Reviser struct {
	bytes []byte
	validation.Validator
	modelVersion    string
	upgradedJsonMap map[string]interface{}
}

// NewReviser returns an upgrader with a validator created from the desired version.
func NewReviser(model utils.InterfaceOrBytes, desiredVersion string) (upgrader Reviser, err error) {
	validator, err := validation.NewValidatorDesiredVersion(model, desiredVersion)
	if err != nil {
		return upgrader, err
	}

	version, err := utils.GetOscalVersionFromMap(validator.GetJsonModel())
	if err != nil {
		return upgrader, err
	}

	return Reviser{
		Validator:       validator,
		modelVersion:    version,
		upgradedJsonMap: nil,
	}, nil
}

// GetModelVersion returns the version of the model being upgraded.
func (u *Reviser) GetModelVersion() string {
	return u.modelVersion
}

// GetRevisedJsonMap sets the upgraded model.
func (u *Reviser) GetRevisedJsonMap() map[string]interface{} {
	return u.upgradedJsonMap
}

func (u *Reviser) Revise() (err error) {
	err = u.Validate()
	if err != nil {
		return err
	}

	originalBytes, err := json.Marshal(u.GetJsonModel())
	if err != nil {
		return err
	}

	upgradedJsonMap, err := utils.CoerceToJsonMap(originalBytes)
	if err != nil {
		return err
	}

	upgradedJsonMap, err = utils.ReplaceOscalVersionInMap(upgradedJsonMap, u.GetSchemaVersion())
	if err != nil {
		return err
	}

	u.upgradedJsonMap = upgradedJsonMap

	return nil
}

// GetRevisedBytes returns the upgraded model as bytes, marshalled to the desired extension.
func (u *Reviser) GetRevisedBytes(ext string) (bytes []byte, err error) {
	switch ext {
	case "json":
		bytes, err = json.Marshal(u.GetRevisedJsonMap())
		if err != nil {
			return nil, err
		}
	case "yaml":
		bytes, err = yaml.Marshal(u.GetRevisedJsonMap())
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}
