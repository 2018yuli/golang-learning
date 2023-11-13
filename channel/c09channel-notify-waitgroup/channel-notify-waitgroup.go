package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received data %c \n", id, n)
		// 通知外部，本函数完成任务，这里需要在另外的 goroutine 中等待
		wg.Done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(i, w.in, wg)
	return w
}

func channelNotify() {
	var workers [10]worker
	var wg sync.WaitGroup

	// consumer
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	// producer
	for i, worker := range workers {
		worker.in <- 'a' + i
		// wg.Add(1)
	}
	// producer2
	for i, worker := range workers {
		worker.in <- 'A' + i
		// wg.Add(1)
	}

	wg.Wait()

}

func main() {
	channelNotify()
}
