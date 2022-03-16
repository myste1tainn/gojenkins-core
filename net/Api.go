package net

import (
	"errors"
	"github.com/myste1tainn/gojenkins-core/helpers"
	"github.com/myste1tainn/gojenkins-core/helpers/logger"
	"github.com/myste1tainn/gojenkins-core/models"
	"io/ioutil"
	"path"
)

type Api interface {
	BuildWithParameters(jobUrl string, params map[string]string) error
	ConsoleText(jobUrl string, jobId string) string
}

type DefaultApi struct {
	Context models.Context
	Http    Http
}

func (this DefaultApi) BuildWithParameters(jobUrl string, params map[string]string) error {
	logger.Log("BuildWithParameters jobUrl:%s params:%s", jobUrl, params)
	uri := path.Join(jobUrl, "buildWithParameters")
	res, err := this.Http.post(uri, params)
	if err != nil {
		return err
	} else if res.StatusCode > 400 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		} else {
			err := errors.New(string(data))
			return err
		}
	} else {
		return nil
	}
}

func (this DefaultApi) ConsoleText(jobUrl string,
	jobId string) string {
	logger.Log("ConsoleText jobUrl:%s jobId:%s", jobUrl, jobId)
	uri := path.Join(jobUrl, jobId, "consoleText")
	resp := helpers.PanicHttpResponse(this.Http.get(uri))
	defer helpers.Panic(resp.Body.Close())
	data := helpers.PanicData(ioutil.ReadAll(resp.Body))
	return string(data)
}
