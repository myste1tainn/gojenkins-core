package UrlUtils

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"strconv"
)

func BaseUrl(m models.Machine) string {
	protocol := "http"
	if m.IsSsl {
		protocol = "https"
	}

	port := ""
	if m.Port > 0 {
		port = ":" + strconv.Itoa(m.Port)
	}

	return protocol + "://" + m.Host + port
}
