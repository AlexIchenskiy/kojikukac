package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Admin struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

func GetAdmins() []Admin {
	yamlFile, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
		return nil
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
		return nil
	}

	var adminList []Admin

	err = yaml.Unmarshal(yamlData, &adminList)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
		return nil
	}

	return adminList
}

var AdminTokens []string
