package utils

import (
	"os"
)

// ReadFile takes a path and returns a byte array and an error
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile takes a path, a byte array, and a file mode and writes the byte array to the specified path
func WriteFile(path string, bytes []byte, perm os.FileMode) error {
	return os.WriteFile(path, bytes, perm)
}

// ConvertToYaml takes a json byte array, coerces it to a yaml []byte, and writes it to the specified path
// func ConvertToYaml(inputPath string, outputPath, string) error {
// 	// model := validation.GetModels(version)
// 	err := yaml.Unmarshal(bytes, &model)
// 	if err != nil {
// 		return err
// 	}

// 	yamlBytes, err := yaml.Marshal(model)
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(outputPath, yamlBytes, 0644)
// }

// ConvertToJSON takes a yaml byte array, coerces it to a json []byte, and writes it to the specified path
// func ConvertToJSON(version string, outputPath string, bytes []byte) error {
// 	model := validation.GetModels(version)
// 	err := yaml.Unmarshal(bytes, &model)
// 	if err != nil {
// 		return err
// 	}

// 	jsonBytes, err := json.Marshal(model)
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(outputPath, jsonBytes, 0644)
// }
