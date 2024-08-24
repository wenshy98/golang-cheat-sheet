package singleton

import (
	"testing"
	"time"
)

func Test_getDoubleCheckInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			if i := getDoubleCheckInstance(); i == nil {
				t.Error("instance is nil")
			}
		}()
	}
	time.Sleep(time.Second)
}
