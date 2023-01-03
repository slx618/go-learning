package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TRACE_CODE string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()

	wg.Add(1)
	key := TRACE_CODE("trace_code")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Println(traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TRACE_CODE("trace_code"), "ssss")
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()

	fmt.Println("Over")

}
