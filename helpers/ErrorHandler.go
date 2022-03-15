package helpers

import "net/http"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicHttpResponse(resp *http.Response, err error) *http.Response {
	if err != nil {
		panic(err)
	} else {
		return resp
	}
}

func PanicData(resp []byte, err error) []byte {
	if err != nil {
		panic(err)
	} else {
		return resp
	}
}
