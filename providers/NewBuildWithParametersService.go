package providers

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/net"
	"github.com/myste1tainn/gojenkins-core/services"
)

func NewBuildWithParametersService(
	context models.Context,
	api net.Api,
) services.BuildWithParametersService {
	return services.DefaultBuildWithParametersService{
		Context: context,
		Api:     api,
	}
}
