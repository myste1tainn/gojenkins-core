package net

import (
	"github.com/myste1tainn/gojenkins-core/models"
)

type Api interface {
	BuildWithParameters(jobUrl string, params map[string]string)
}

type DefaultApi struct {
	Context models.Context
	Http    Http
}

func (this DefaultApi) BuildWithParameters(jobUrl string, params map[string]string) {
	_, err := this.Http.post(jobUrl, params)
	if err != nil {
		panic(err)
	}
}
