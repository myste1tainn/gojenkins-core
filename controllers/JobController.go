package controllers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/helpers/logger"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/services"
)

type JobController interface {
	BuildWithParameters(req requests.BuildWithParametersRequest) error
	ConsoleText(req requests.ConsoleTextRequest) string
}

type DefaultJobController struct {
	ConfigLoader               config.MainYamlConfigLoader
	BuildWithParametersService services.BuildWithParametersService
	ConsoleTextService         services.ConsoleTextService
}

func (d DefaultJobController) BuildWithParameters(req requests.BuildWithParametersRequest) error {
	logger.Log("BuildWithParameters %v", req)
	return d.BuildWithParametersService.Execute(req)
}

func (d DefaultJobController) ConsoleText(req requests.ConsoleTextRequest) string {
	logger.Log("ConsoleText %v", req)
	return d.ConsoleTextService.Execute(req)
}
