package providers

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/net"
)

func NewApi(
	ctx models.Context,
	http net.Http,
) net.Api {
	return net.DefaultApi{
		Context: ctx,
		Http:    http,
	}
}
