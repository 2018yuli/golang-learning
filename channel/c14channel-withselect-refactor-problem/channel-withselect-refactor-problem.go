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
		time.Sleep(5 * time.Second)
		fmt.Printf("worker %d received data %d \n", id, n)
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

func SelectRefactorProblem() {
	// select 的使用场景，我们有多个 channel
	var c1, c2 = Generate(), Generate()

	var worker = createWorker(0)
	n := 0
	hasValue := false
	for i := 0; i < 50; i++ {
		// nil channel 能正确运行，但是不会被 select 到
		var activeWorker chan int
		if hasValue {
			activeWorker = worker.in
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
			// 但是，如果 c1, c2 生产数据的速度大于 activeWorker 消费数据的速度
			// n 将会无法接收，所以此程序依旧有问题
		case activeWorker <- n:
			hasValue = false
		}
	}

}

func main() {
	SelectRefactorProblem()
}
