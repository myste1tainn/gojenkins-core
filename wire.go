//+build wireinject

package gojenkins_core

import (
	"github.com/google/wire"
	"github.com/myste1tainn/gojenkins-core/controllers"
	"github.com/myste1tainn/gojenkins-core/providers"
)

func InitializeJobController() controllers.JobController {
	wire.Build(
		providers.NewMainYamlConfigLoader,
		providers.NewFileReader,
		providers.NewApi,
		providers.NewHttp,
		providers.NewContext,
		providers.NewBuildWithParametersService,
		providers.NewJobController,
	)
	return nil
}
