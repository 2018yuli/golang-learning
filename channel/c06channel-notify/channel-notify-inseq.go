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
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}
	// producer2
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}

}

func main() {
	channelNotify()
}
