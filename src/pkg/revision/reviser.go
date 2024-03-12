package revision

import (
	"encoding/json"
	"time"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/goccy/go-yaml"
)

type Reviser struct {
	bytes []byte
	validation.Validator
	modelVersion    string
	upgradedJsonMap yaml.MapSlice
}

// NewReviser returns an upgrader with a validator created from the desired version.
func NewReviser(model utils.InterfaceOrBytes, desiredVersion string) (upgrader Reviser, err error) {
	ogBytes, ok := model.([]byte)
	if ok {
		upgrader.bytes = ogBytes
	}

	validator, err := validation.NewValidatorDesiredVersion(model, desiredVersion)
	if err != nil {
		return upgrader, err
	}
	upgrader.Validator = validator

	version, err := utils.GetOscalVersionFromMap(validator.GetJsonModel())
	if err != nil {
		return upgrader, err
	}
	upgrader.modelVersion = version

	return upgrader, nil
}

// GetModelVersion returns the version of the model being upgraded.
func (u *Reviser) GetModelVersion() string {
	return u.modelVersion
}

// GetRevisedJsonMap sets the upgraded model.
func (u *Reviser) GetRevisedJsonMap() yaml.MapSlice {
	return u.upgradedJsonMap
}

func (u *Reviser) Revise() (err error) {
	err = u.Validate()
	if err != nil {
		return err
	}

	var originalBytes []byte
	if u.bytes != nil {
		originalBytes = u.bytes
	} else {
		originalBytes, err = json.Marshal(u.GetJsonModel())
		if err != nil {
			return err
		}
	}

	var upgradedJsonMap yaml.MapSlice // UseOrderedMap() is used to preserve the order of the keys in the map
	err = yaml.UnmarshalWithOptions(originalBytes, &upgradedJsonMap, yaml.UseOrderedMap())
	if err != nil {
		return err
	}

	replacedValue, err := utils.ReplaceKeyInMapSlice(upgradedJsonMap[0].Value.(yaml.MapSlice), []string{"metadata", "oscal-version"}, u.GetSchemaVersion())
	if err != nil {
		return err
	}
	upgradedJsonMap[0].Value = replacedValue

	replacedValue, err = utils.ReplaceKeyInMapSlice(upgradedJsonMap[0].Value.(yaml.MapSlice), []string{"metadata", "last-modified"}, time.Now())
	if err != nil {
		return err
	}
	upgradedJsonMap[0].Value = replacedValue

	u.upgradedJsonMap = upgradedJsonMap

	return nil
}

// GetRevisedBytes returns the upgraded model as bytes, marshalled to the desired extension.
func (u *Reviser) GetRevisedBytes(ext string) (bytes []byte, err error) {
	return utils.MarshalByExtension(u.upgradedJsonMap, ext)
}
