package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received data %c \n", id, n)
		// 通知外部，本函数完成任务
		// 因为 goroutine 的发送数据是阻塞的，当第二个 producer 过来发数据时
		// done 中的channel 因为没有 goroutine 接收，从而引发死锁
		done <- true
	}
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(i, w.in, w.done)
	return w
}

func channelNotify() {
	var workers [10]worker

	// consumer
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// producer
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// producer2
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// waiting for all task to complete
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}

}

func main() {
	channelNotify()
}
