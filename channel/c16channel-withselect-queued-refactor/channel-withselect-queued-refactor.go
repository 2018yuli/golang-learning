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
		time.Sleep(1 * time.Second)
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

func SelectQueuedRefactor() {
	// select 的使用场景，我们有多个 channel
	var c1, c2 = Generate(), Generate()

	var worker = createWorker(0)
	var queue []int
	tm := time.After(30 * time.Second)
	tk := time.Tick(time.Second)
	for {
		// nil channel 能正确运行，但是不会被 select 到
		var activeWorker chan int
		var activeValue int
		if len(queue) > 0 {
			activeWorker = worker.in
			activeValue = queue[0]
		}
		select {
		case n := <-c1:
			queue = append(queue, n)
		case n := <-c2:
			queue = append(queue, n)
		case activeWorker <- activeValue:
			queue = queue[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out after 800ms")
		case <-tm:
			fmt.Println("Bye!")
			return
		case <-tk:
			fmt.Println("current queue length ================== :", len(queue))
		}
	}

}

func main() {
	SelectQueuedRefactor()
}
