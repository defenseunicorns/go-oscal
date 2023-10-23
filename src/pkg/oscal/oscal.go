package oscal

import (
	"log"

	"gopkg.in/yaml.v3"
)

// Keep this generic for possible use with all models
// Perform identification of model and current version
// Pass that data to functions that perform specific updates
func ConvertOscal(data []byte, targetVersion string) (any, error) {
	temp := make(map[string]interface{})
	var currentVersion string

	err := yaml.Unmarshal(data, &temp)
	if err != nil {
		return nil, err
	}
	if val, ok := temp["system-security-plan"]; ok {
		log.Default().Println("System Security Plan Detected")
		value, ok := val.(map[string]interface{})["metadata"].(map[string]interface{})["oscal-version"]
		if ok {
			currentVersion = value.(string)
			log.Default().Println("currentVersion: ", currentVersion)
		}
	} else if val, ok := temp["component-definition"]; ok {
		log.Default().Println("Component Definition Detected")
		value, ok := val.(map[string]interface{})["metadata"].(map[string]interface{})["oscal-version"]
		if ok {
			currentVersion = value.(string)
			log.Default().Println("currentVersion: ", currentVersion)
			ConvertComponentDefinition(data, currentVersion, targetVersion)
		}
	} else {
		log.Fatal("OSCAL model not supported")
	}

	// log.Default().Println("temp: ", temp["component-definition"])

	return nil, nil
}
