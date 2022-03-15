package helpers

import "io/ioutil"

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type DefaultFileReader struct {
}

func (d DefaultFileReader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
