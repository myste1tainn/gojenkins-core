package providers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/models"
)

func NewContext(loader config.MainYamlConfigLoader) models.Context {
	c := loader.Load()

	serverConfig := c.Core.Server
	isSsl := false
	if serverConfig.IsSsl == "true" {
		isSsl = true
	}
	host := models.Machine{
		IsSsl: isSsl,
		Host:  serverConfig.Host,
		Port:  serverConfig.Port,
	}

	userConfig := c.Core.User
	user := models.User{
		Username: userConfig.Username,
		Token:    userConfig.Token,
	}

	return models.Context{
		Host: host,
		User: user,
	}
}
