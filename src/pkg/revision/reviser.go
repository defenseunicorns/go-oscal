package revision

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
	"gopkg.in/yaml.v3"
)

type Reviser struct {
	bytes []byte
	validation.Validator
	modelVersion    string
	upgradedJsonMap map[string]interface{}
}

// NewReviser returns an upgrader with a validator created from the desired version.
func NewReviser(interfaceOrBytes model.InterfaceOrBytes, desiredVersion string) (upgrader Reviser, err error) {
	validator, err := validation.NewValidatorDesiredVersion(interfaceOrBytes, desiredVersion)
	if err != nil {
		return upgrader, err
	}

	version, err := versioning.GetOscalVersionFromMap(validator.GetJsonModel())
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

	upgradedJsonMap, err := model.CoerceToJsonMap(originalBytes)
	if err != nil {
		return err
	}

	upgradedJsonMap, err = versioning.ReplaceOscalVersionInMap(upgradedJsonMap, u.GetSchemaVersion())
	if err != nil {
		return err
	}

	u.upgradedJsonMap = upgradedJsonMap

	return nil
}

// GetRevisedBytes returns the upgraded model as bytes, marshalled to the desired extension.
func (u *Reviser) GetRevisedBytes(ext string) (bytes []byte, err error) {
	// return model.MarshalByExtension(u.upgradedJsonMap, ext)
	split := strings.Split(ext, ".")
	ext = split[len(split)-1]

	switch ext {
	case "json":
		bytes, err = json.Marshal(u.upgradedJsonMap)
		if err != nil {
			return nil, err
		}

	case "yaml":
		bytes, err = yaml.Marshal(u.upgradedJsonMap)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	// Use the intended upgrade schema version
	switch u.GetSchemaVersion() {
	case "1.0.4":
		return model.Marshall104(bytes, ext)
	case "1.0.6":
		return model.Marshall106(bytes, ext)
	case "1.1.0":
		return model.Marshall110(bytes, ext)
	case "1.1.1":
		return model.Marshall111(bytes, ext)
	case "1.1.2":
		return model.Marshall112(bytes, ext)
	default:
		return nil, fmt.Errorf("unsupported model version")
	}

}
