package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f2(ctx context.Context) {
	wg.Add(1)
	defer wg.Done()
breakTag:
	for {
		fmt.Println("demo2")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break breakTag
		default:
		}
	}

}

func f(ctx context.Context) {
	wg.Add(1)
	go f2(ctx)
	defer wg.Done()
breakTag:
	for {
		fmt.Println("demo")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break breakTag
		default:
		}
	}

}

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
