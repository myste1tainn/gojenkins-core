package providers

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/net"
)

func NewHttp(
	ctx models.Context,
) net.Http {
	sender := net.DefaultHttpRequestSender{
		Context: ctx,
	}
	return net.DefaultHttp{
		Context: ctx,
		Sender:  sender,
	}
}
