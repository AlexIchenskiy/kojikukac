package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var fileCnf *Configurations = nil

func GetConfigFromYaml() (*Configurations, error) {
	if fileCnf == nil {
		yamlFile, err := os.ReadFile("config/config.yaml")
		if err != nil {
			return nil, err
		}
		temp := &Configurations{}
		err = yaml.Unmarshal(yamlFile, temp)
		if err != nil {
			return nil, err
		}
		fileCnf = temp

	}
	return fileCnf, nil
}

type Configurations struct {
	JmbagConf string   `yaml:"jmbag"`
	HTTPConf  httpConf `yaml:"http"`
	UsersConf []user   `yaml:"users"`
}

type httpConf struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type user struct {
	Name     string `yaml:"name"`
	Jmbag    string `yaml:"jmbag"`
	Password string `yaml:"password"`
}
