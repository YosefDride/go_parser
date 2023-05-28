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


func parse(filePath string, key string, replaceString string) error {
	data, err := readYAMLFile(filePath)
	if err != nil {
		return err
	}

	// Check if the key exists in the data
	if _, ok := data[key]; !ok {
		return fmt.Errorf("key not found: %s", key)
	}

	// Update the value for the key with the replaceString
	data[key] = replaceString

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
	key := "paths.cam1.source"
	replaceString := "NEWIP"
	time.Sleep(2 * time.Second)

	err := parse(filePath, key, replaceString)
	fmt.Println(err)
	time.Sleep(2 * time.Second)

	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	fmt.Println("YAML file parsed successfully.")
	time.Sleep(5 * time.Second)
}