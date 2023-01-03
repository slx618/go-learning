package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(5 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), t)

	//尽量手动调用一下 cancel
	//不这样做 可能会使上下文及其父类没有释放
	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println("delay 1 sec")
	case <-ctx.Done():
		fmt.Println("done")
	}
}
