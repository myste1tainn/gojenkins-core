package net

import (
	"bytes"
	"github.com/myste1tainn/gojenkins-core/helpers/UrlUtils"
	"github.com/myste1tainn/gojenkins-core/helpers/logger"
	"github.com/myste1tainn/gojenkins-core/models"
	"mime/multipart"
	"net/http"
)

type Http interface {
	get(uri string) (*http.Response, error)
	post(uri string, params map[string]string) (*http.Response, error)
}

type DefaultHttp struct {
	Context models.Context
	Sender  HttpRequestSender
}

func (this DefaultHttp) get(uri string) (*http.Response, error) {
	fullUrl := UrlUtils.FullUrl(this.Context.Host, uri)
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	return this.performRequest(req, err)
}

func (this DefaultHttp) post(uri string, params map[string]string) (*http.Response, error) {
	fullUrl := UrlUtils.FullUrl(this.Context.Host, uri)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		err := writer.WriteField(key, val)
		if err != nil {
			panic(err)
		}
	}
	_ = writer.Close()
	req, err := http.NewRequest(http.MethodPost, fullUrl, bytes.NewReader(body.Bytes()))
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return this.performRequest(req, err)
}

func (this DefaultHttp) performRequest(req *http.Request, err error) (*http.Response, error) {
	logger.Log("Performing Requests to %s %s\n", req.Method, req.URL)
	if err != nil {
		panic(err)
	} else {
		return this.Sender.Send(req)
	}
}
