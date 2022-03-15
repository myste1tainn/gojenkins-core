package tests

import (
	"fmt"
	"github.com/myste1tainn/gojenkins-core/config"
	"github.com/myste1tainn/gojenkins-core/tests/mocks"
	"testing"
)

func TestLoad_givenCorrectStructureYaml_itReturnsConfig(t *testing.T) {
	// given
	isSsl := "true"
	host := "localhost"
	port := 443
	username := "test_username"
	token := "test_token"
	content := fmt.Sprintf(`
core:
  server: 
    ssl: %s
    host: %s
    port: %d
  user: 
    username: %s
    token: %s
`, isSsl, host, port, username, token)
	fileReader := mocks.MockFileReader{
		FileContent: content,
	}

	// then
	mConfig := config.DefaultMainYamlConfigLoader{
		FileReader: fileReader,
	}.Load()

	// expect
	server := mConfig.Core.Server
	if server.IsSsl != isSsl {
		t.Errorf("Expected isSsl %s, got %s", isSsl, server.IsSsl)
	}

	if server.Host != host {
		t.Errorf("Expected host %s, got %s", host, server.Host)
	}

	if server.Port != port {
		t.Errorf("Expected port %d, got %d", port, server.Port)
	}

	user := mConfig.Core.User
	if user.Username != username {
		t.Errorf("Expected username %s, got %s", username, user.Username)
	}

	if user.Token != token {
		t.Errorf("Expected token %s, got %s", token, user.Token)
	}
}
