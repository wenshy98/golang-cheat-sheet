package usecontext

import (
	"context"
	"testing"
	"time"
)

func Test_cancelSpeak(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go cancelSpeak(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
