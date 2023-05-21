package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"./Config"
	"gopkg.in/yaml.v2"
)

// This function readYAMLFile takes a filePath string as input and returns a map[string]interface{} representing the YAML data and an error if any occurred.
// It uses ioutil.ReadFile to read the contents of the file specified by filePath and stores it in yamlFile.
// Then, it declares a variable data as a map[string]interface{} to hold the unmarshaled YAML data.
// It uses yaml.Unmarshal to convert the YAML data in yamlFile into the data map.
// If any error occurred during reading or unmarshaling, it returns nil for data and the error.

func readYAMLFile(filePath string) (map[string]interface{}, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// This function updateFieldValue takes a data map, a key string, and a newValue interface{} as input and returns an error if any occurred.
// It checks if the key exists in the data map by using a comma-ok idiom (_, ok := data[key]). If the key is not found, it returns an error using fmt.Errorf.
// If the key exists, it updates the value for that key in the data map with the provided newValue.

func updateFieldValue(filePath string, key string, newValue interface{}) error {
	data, err := readYAMLFile(filePath)
	if err != nil {
		return err
	}

	// Check if the key exists in the data
	if _, ok := data[key]; !ok {
		return fmt.Errorf("key not found: %s", key)
	}

	// Update the value for the key
	data[key] = newValue

	err = writeYAMLFile(filePath, data)
	if err != nil {
		return err
	}

	return nil
}

// This function writeYAMLFile takes a filePath string and a data map as input and returns an error if any occurred.
// It uses yaml.Marshal to convert the data map into YAML format and assigns the resulting YAML data to yamlData.
// It uses ioutil.WriteFile to write the yamlData to the file specified by filePath.
// The file is written with file permissions 0644, which grants read and write permissions to the owner and read-only permissions to others.
// If any error occurred during marshaling or writing, it returns the error.

func writeYAMLFile(filePath string, data map[string]interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {

    filePath := Config.FilePath

	err := updateFieldValue(filePath, "field3", "noob")
	if err != nil {
		log.Fatalf("Failed to update field value: %v", err)
	}

	fmt.Println("YAML file modified successfully.")
}