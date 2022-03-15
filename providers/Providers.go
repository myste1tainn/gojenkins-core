package providers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/controllers"
	"github.com/myste1tainn/gojenkins-core/helpers"
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/net"
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

func NewMainYamlConfigLoader(reader helpers.FileReader) config.MainYamlConfigLoader {
	return config.DefaultMainYamlConfigLoader{
		FileReader: reader,
	}
}

func NewFileReader() helpers.FileReader {
	return helpers.DefaultFileReader{}
}

func NewBuildWithParametersService(
	context models.Context,
	api net.Api,
) services.BuildWithParametersService {
	return services.DefaultBuildWithParametersService{
		Context: context,
		Api:     api,
	}
}

func NewApi(
	ctx models.Context,
	http net.Http,
) net.Api {
	return net.DefaultApi{
		Context: ctx,
		Http:    http,
	}
}

func NewHttp(
	ctx models.Context,
	sender net.HttpRequestSender,
) net.Http {
	return net.DefaultHttp{
		Context: ctx,
		Sender:  sender,
	}
}
