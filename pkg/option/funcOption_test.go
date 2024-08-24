package option

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer(WithHost("127.0.0.1:8080"), WithPort(8080))
	if server.Host != "127.0.0.1:8080" {
		t.Errorf("server host is incorrect %s", server.Host)
	}
	if server.Port != 8080 {
		t.Errorf("server port is incorrect %d", server.Port)
	}
}
