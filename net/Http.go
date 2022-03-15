package net

import (
	"github.com/myste1tainn/gojenkins-core/helpers/UrlUtils"
	"github.com/myste1tainn/gojenkins-core/models"
	"net/http"
	"net/url"
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
	req, err := http.NewRequest(http.MethodPost, fullUrl, nil)
	form := url.Values{}
	for key, value := range params {
		form.Add(key, value)
	}
	req.PostForm = form
	return this.performRequest(req, err)
}

func (this DefaultHttp) performRequest(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		panic(err)
	} else {
		return this.Sender.Send(req)
	}
}
