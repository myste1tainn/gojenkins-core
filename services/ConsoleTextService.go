package services

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/net"
)

type ConsoleTextService interface {
	Execute(req requests.ConsoleTextRequest) string
}

type DefaultConsoleTextService struct {
	Context models.Context
	Api     net.Api
}

func (this DefaultConsoleTextService) Execute(req requests.ConsoleTextRequest) string {
	return this.Api.ConsoleText(req.JobUrl, req.JobId)
}
