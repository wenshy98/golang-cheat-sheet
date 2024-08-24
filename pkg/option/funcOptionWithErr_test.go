package option

import (
	"testing"
)

func TestNewServer2(t *testing.T) {
	server, err := NewServer2(WithHost2("127.0.0.1:8080"), WithPort2(8080))
	if server == nil && err != nil {
		t.Errorf("err should be nil but got %v", err.Error())
	}
	if server.Host != "127.0.0.1:8080" {
		t.Errorf("server host is incorrect %s", server.Host)
	}
	if server.Port != 8080 {
		t.Errorf("server port is incorrect %d", server.Port)
	}
	server, err = NewServer2(WithHost2("127.0.0.1:8080"), WithPort2(0))
	if err == nil {
		t.Errorf("err should not be nil")
	}
}
