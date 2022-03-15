package controllers

import (
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/net"
	"github.com/myste1tainn/gojenkins-core/services"
)

type JobController interface {
	BuildWithParameters(req requests.BuildWithParametersRequest)
}

type DefaultJobController struct {
	ConfigLoader config.MainYamlConfigLoader
}

func (d DefaultJobController) BuildWithParameters(req requests.BuildWithParametersRequest) {
	service := services.DefaultBuildWithParametersService{
		Context: context,
		Api: net.DefaultApi{
			Context: context,
			Http: net.DefaultHttp{
				Context: context,
				Sender:  net.DefaultHttpRequestSender{Context: context},
			},
		},
	}
}