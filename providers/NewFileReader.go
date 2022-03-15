package providers

import "github.com/myste1tainn/gojenkins-core/helpers"

func NewFileReader() helpers.FileReader {
	return helpers.DefaultFileReader{}
}
