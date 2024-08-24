package option

import "testing"

func TestDoSomething(t *testing.T) {
	err := DoSomething("1", WithValue(1))
	if err != nil {
		t.Errorf("error should be nil")
	}
}
