package singleton

import (
	"testing"
	"time"
)

func Test_getInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			instance := getInstance()
			if instance == nil {
				t.Error("instance is nil")
			}
		}()
	}
	time.Sleep(time.Second)
}

func Test_getDoubleCheckInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			instance := getDoubleCheckInstance()
			if instance == nil {
				t.Error("instance is nil")
			}
		}()
	}
	time.Sleep(time.Second)
}
