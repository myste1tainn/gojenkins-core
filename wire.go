package main

import (
	"github.com/google/wire"
	"github.com/myste1tainn/gojenkins-core/controllers"
	"github.com/myste1tainn/gojenkins-core/providers"
)

func InitializeJobController() controllers.JobController {
	wire.Build(
		providers.NewJobController,
		providers.NewMainYamlConfigLoader,
		providers.NewFileReader,
		providers.NewBuildWithParametersService,
		providers.NewApi,
		providers.NewHttp,
	)
	return InitializeJobController()
}
