package config

import (
	"github.com/myste1tainn/gojenkins-core/helpers"
	"gopkg.in/yaml.v2"
)

var MainYamlConfigFileName = "gojenkins.yaml"

type MainYamlConfigLoader interface {
	Load() MainYamlConfig
}

type DefaultMainYamlConfigLoader struct {
	FileReader helpers.FileReader
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
