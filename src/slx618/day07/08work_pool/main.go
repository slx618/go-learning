package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type rst struct {
	*job
	rst int64
	id  int
}

var jobChan = make(chan *job, 100)
var rstChan = make(chan *rst, 100)

func sum(i int64) (rst int64) {
	for i > 0 {
		rst += i % 10
		i = i / 10
	}
	return
}

var wg sync.WaitGroup

func addJob(ch chan<- *job) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	for {
		ch <- &job{
			value: rand.Int63(),
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func work(ch1 <-chan *job, ch2 chan<- *rst, i int) {
	defer wg.Done()
	for {
		ch := <-ch1
		//if !ok {
		//	continue
		//}
		//fmt.Println("job start, job value: ", ch.value)
		ch2 <- &rst{
			ch,
			sum(ch.value),
			i,
		}
		//fmt.Println("job value counted")
	}

}

func main() {
	wg.Add(1)
	go addJob(jobChan)
	for i := 0; i <= 24; i++ {
		wg.Add(1)
		go work(jobChan, rstChan, i)
	}

	for ch := range rstChan {
		fmt.Println("work id:", ch.id, "value: ", ch.job.value, "rst: ", ch.rst)
	}
	wg.Wait()

}
