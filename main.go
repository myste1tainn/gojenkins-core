package main

import (
	"github.com/myste1tainn/gojenkins-core/models/requests"
	gojenkins_core "github.com/myste1tainn/gojenkins-core/wire"
)

func main() {
	ctrl := gojenkins_core.InitializeJobController()
	ctrl.BuildWithParameters(
		requests.BuildWithParametersRequest{
			JobUrl: "job/test/job/test-hell-world",
			Params: nil,
		},
	)
}
