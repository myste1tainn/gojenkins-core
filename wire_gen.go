// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package gojenkins_core

import (
	"github.com/myste1tainn/gojenkins-core/controllers"
	"github.com/myste1tainn/gojenkins-core/providers"
)

// Injectors from wire.go:

func InitializeJobController() controllers.JobController {
	fileReader := providers.NewFileReader()
	mainYamlConfigLoader := providers.NewMainYamlConfigLoader(fileReader)
	context := providers.NewContext(mainYamlConfigLoader)
	http := providers.NewHttp(context)
	api := providers.NewApi(context, http)
	buildWithParametersService := providers.NewBuildWithParametersService(context, api)
	jobController := providers.NewJobController(mainYamlConfigLoader, buildWithParametersService)
	return jobController
}