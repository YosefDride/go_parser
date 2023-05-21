package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"time"
	"parser/config"
)

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
	fmt.Println(filePath)
	err := updateFieldValue(filePath, "field3", "damn")
	if err != nil {
		log.Fatalf("Failed to update field value: %v", err)
	}
	fmt.Println("YAML file modified successfully.")

	time.Sleep(10 * time.Second)
}