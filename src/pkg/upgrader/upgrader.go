package upgrader

import (
	"encoding/json"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

type Upgrader struct {
	bytes []byte
	validation.Validator
	modelVersion    string
	upgradedJsonMap map[string]interface{}
}

// NewUpgrader returns an upgrader with a validator created from the desired version.
func NewUpgrader(model utils.InterfaceOrBytes, desiredVersion string) (upgrader Upgrader, err error) {
	validator, err := validation.NewValidatorDesiredVersion(model, desiredVersion)
	if err != nil {
		return upgrader, err
	}

	version, err := utils.GetOscalVersionFromMap(validator.GetJsonModel())
	if err != nil {
		return upgrader, err
	}

	return Upgrader{
		Validator:       validator,
		modelVersion:    version,
		upgradedJsonMap: nil,
	}, nil
}

// GetModelVersion returns the version of the model being upgraded.
func (u *Upgrader) GetModelVersion() string {
	return u.modelVersion
}

// GetUpgradedJsonMap sets the upgraded model.
func (u *Upgrader) GetUpgradedJsonMap() map[string]interface{} {
	return u.upgradedJsonMap
}

func (u *Upgrader) Upgrade() (err error) {
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
