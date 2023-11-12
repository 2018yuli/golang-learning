package main

import (
	"fmt"
	"time"
)

func GorutineFuncWithChanParam(c <-chan int) {
	for {
		n := <-c
		fmt.Printf("n= %v \n", n)
	}
}

func ChannelAsParam() {
	c := make(chan int)
	go GorutineFuncWithChanParam(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func GorutineWorker(i int, c <-chan int) {
	for {
		fmt.Printf("worker %d received data %c \n", i, <-c)
	}
}

func SimpleConsumerProducerDemo() {

	var channels [10]chan int

	// consumer
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go GorutineWorker(i, channels[i])
	}

	// producer
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	// producer2
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	SimpleConsumerProducerDemo()
}
