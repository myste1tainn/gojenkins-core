package net

import (
	"github.com/myste1tainn/gojenkins-core/helpers"
	"github.com/myste1tainn/gojenkins-core/models"
	"io"
	"path"
)

type Api interface {
	BuildWithParameters(jobUrl string, params map[string]string)
	ConsoleText(jobUrl string, jobId string) string
}

type DefaultApi struct {
	Context models.Context
	Http    Http
}

func (this DefaultApi) BuildWithParameters(jobUrl string, params map[string]string) {
	uri := path.Join(jobUrl, "buildWithParameters")
	_ = helpers.PanicHttpResponse(this.Http.post(uri, params))
}

func (this DefaultApi) ConsoleText(jobUrl string, jobId string) string {
	uri := path.Join(jobUrl, jobId, "consoleText")
	resp := helpers.PanicHttpResponse(this.Http.get(uri))
	defer helpers.Panic(resp.Body.Close())
	data := helpers.PanicData(io.ReadAll(resp.Body))
	return string(data)
}
