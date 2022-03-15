package services

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/net"
)

type BuildWithParametersService interface {
	Execute(req requests.BuildWithParametersRequest)
}

type DefaultBuildWithParametersService struct {
	Context models.Context
	Api     net.Api
}

func (this DefaultBuildWithParametersService) Execute(req requests.BuildWithParametersRequest) {
	this.Api.BuildWithParameters(req.JobUrl, req.Params)
}
