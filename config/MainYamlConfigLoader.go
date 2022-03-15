package config

import (
	"gopkg.in/yaml.v2"
)

var MainYamlConfigFileName = "gojenkins.yaml"

type MainYamlConfigLoader interface {
	Load() MainYamlConfig
}

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type DefaultMainYamlConfigLoader struct {
	FileReader FileReader
}

func (this DefaultMainYamlConfigLoader) Load() MainYamlConfig {
	data, err := this.FileReader.ReadFile("./" + MainYamlConfigFileName)
	if err != nil {
		panic(err)
	} else {
		var config MainYamlConfig
		err := yaml.Unmarshal(data, &config)
		if err != nil {
			panic(err)
		}
		return config
	}
}
