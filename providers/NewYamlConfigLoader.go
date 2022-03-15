package providers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/helpers"
)

func NewMainYamlConfigLoader(reader helpers.FileReader) config.MainYamlConfigLoader {
	return config.DefaultMainYamlConfigLoader{
		FileReader: reader,
	}
}
