package net

import (
	"encoding/base64"
	"github.com/myste1tainn/gojenkins-core/models"
	"net/http"
)

type HttpRequestSender interface {
	Send(req *http.Request) (*http.Response, error)
}

type DefaultHttpRequestSender struct {
	Context models.Context
}

func (this DefaultHttpRequestSender) Send(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", AuthorizationValue(this.Context.User))
	client := http.Client{}
	return client.Do(req)
}

func AuthorizationValue(u models.User) string {
	credentials := u.Username + ":" + u.Token
	encodedString := base64.StdEncoding.EncodeToString([]byte(credentials))
	return "Basic " + encodedString
}
