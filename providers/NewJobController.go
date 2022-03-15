package providers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/controllers"
	"github.com/myste1tainn/gojenkins-core/services"
)

func NewJobController(
	configLoader config.MainYamlConfigLoader,
	bwpService services.BuildWithParametersService,
) controllers.JobController {
	return controllers.DefaultJobController{
		ConfigLoader:               configLoader,
		BuildWithParametersService: bwpService,
	}
}
