package builder

import (
	"testing"
)

func TestNewConfigBuilder(t *testing.T) {
	config := NewConfigBuilder().ValueA("1").ValueB("2").Build()
	if config.ValueA != "1" {
		t.Errorf("Value A was wrong")
	}
	if config.ValueB != "2" {
		t.Errorf("Value B was wrong")
	}
}
