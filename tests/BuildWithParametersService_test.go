package tests

import (
	"github.com/myste1tainn/gojenkins-core/models"
	"github.com/myste1tainn/gojenkins-core/models/requests"
	"github.com/myste1tainn/gojenkins-core/net"
	"github.com/myste1tainn/gojenkins-core/services"
	"github.com/myste1tainn/gojenkins-core/tests/mocks"
	"strconv"
	"testing"
)

func TestBuildWithParameters_givenAllOkContext_callToApiBuildWithParameters(t *testing.T) {
	user := models.User{
		Username: "TestUsername",
		Token:    "TestToken",
	}

	machine := models.Machine{
		IsSsl: true,
		Host:  "localhost",
		Port:  443,
	}

	context := models.Context{
		Host: machine,
		User: user,
	}

	sender := mocks.SpyRequestSender{}

	http := net.DefaultHttp{
		Context: context,
		Sender:  &sender,
	}

	api := net.DefaultApi{
		Context: context,
		Http:    http,
	}

	svc := services.DefaultBuildWithParametersService{
		Context: context,
		Api:     api,
	}

	jobUrl := "job/testJob1/job/testJob2"
	svc.Execute(requests.BuildWithParametersRequest{
		JobUrl: jobUrl,
		Params: map[string]string{
			"test_key": "test_value",
		},
	})

	protocol := "http"
	if machine.IsSsl {
		protocol = "https"
	}
	expected := protocol + "://" + machine.Host + ":" + strconv.Itoa(machine.Port) + "/" + jobUrl
	if sender.SentUrl != expected {
		t.Errorf("Expected %s, got %s", expected, sender.SentUrl)
	}
}
