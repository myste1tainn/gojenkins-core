package controllers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/services"
)

type JobController interface {
	BuildWithParameters(req requests.BuildWithParametersRequest)
	ConsoleText(req requests.ConsoleTextRequest) string
}

type DefaultJobController struct {
	ConfigLoader               config.MainYamlConfigLoader
	BuildWithParametersService services.BuildWithParametersService
	ConsoleTextService         services.ConsoleTextService
}

func (d DefaultJobController) BuildWithParameters(req requests.BuildWithParametersRequest) {
	d.BuildWithParametersService.Execute(req)
}

func (d DefaultJobController) ConsoleText(req requests.ConsoleTextRequest) string {
	return d.ConsoleTextService.Execute(req)
}
