package controllers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/services"
)

type JobController interface {
	BuildWithParameters(req requests.BuildWithParametersRequest)
}

type DefaultJobController struct {
	ConfigLoader               config.MainYamlConfigLoader
	BuildWithParametersService services.BuildWithParametersService
}

func (d DefaultJobController) BuildWithParameters(req requests.BuildWithParametersRequest) {
}
