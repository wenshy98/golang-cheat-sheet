package usecontext

import (
	"context"
	"fmt"
	"time"
)

func cancelSpeak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("Speak, 正在执行中")
		}
	}
}
