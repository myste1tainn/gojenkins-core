package mocks

import "net/http"

type SpyRequestSender struct {
	SentUrl string
}

func (this *SpyRequestSender) Send(req *http.Request) (*http.Response, error) {
	this.SentUrl = req.URL.String()
	return nil, nil
}
