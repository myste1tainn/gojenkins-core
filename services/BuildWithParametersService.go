package services

import (
	"github.com/myste1tainn/gojenkins-core/helpers/logger"
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/net"
)

type BuildWithParametersService interface {
	Execute(req requests.BuildWithParametersRequest) error
}

type DefaultBuildWithParametersService struct {
	Context models.Context
	Api     net.Api
}

func (this DefaultBuildWithParametersService) Execute(req requests.BuildWithParametersRequest) error {
	logger.Log("%v", req)
	return this.Api.BuildWithParameters(req.JobUrl, req.Params)
}
