package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received data %c \n", id, n)
		// 通知外部，本函数完成任务，这里需要在另外的 goroutine 中等待
		w.done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(i, w)
	return w
}

func channelNotify() {
	var workers [10]worker
	var wg sync.WaitGroup

	// consumer
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	// wg.Add(20)

	// producer
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}
	// producer2
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func main() {
	channelNotify()
}
