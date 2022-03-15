package UrlUtils

import "github.com/myste1tainn/gojenkins-core/models"

func FullUrl(m models.Machine, path string) string {
	baseUrl := BaseUrl(m)
	return baseUrl + "/" + path
}
