package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received data %d \n", id, n)
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

func SelectOverview() {
	// select 的使用场景，我们有多个 channel
	var c1, c2 = Generate(), Generate()
	// 我们一个 worker
	w := createWorker(1)
	// 当我们想从多个 channel 中收取数据，谁快用谁，这时，可以使用select
	for i := 0; i < 50; i++ {
		select {
		case n := <-c1:
			// fmt.Printf("\n received value from c1 %d \n", n)
			// 但是这里 w.in <- n 是阻塞的？
			/*
				The statement n <- c is trying to receive a value from channel c.
				If there's no value available on channel c,
				this operation will wait until a value is sent on that channel.
				Until then, the execution of the program will be blocked at this point.
			*/
			w.in <- n
			<-w.done
		case n2 := <-c2:
			// fmt.Printf("\n received value from c2 %d \n", n2)
			w.in <- n2
			<-w.done
			// 加了 defalut 之后 select 将变为非阻塞
			// default:
			// 	fmt.Print(" ... ")
			// 	time.Sleep(time.Duration(7*100) * time.Millisecond)
			// }
		}
	}

}

func main() {
	SelectOverview()
}
