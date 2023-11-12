package main

import (
	"fmt"
	"time"
)

// <-chan receive only, chan<- send only
func GorutineWorker(i int) chan<- int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("worker %d received data %c \n", i, <-c)
		}
	}()

	return c
}

func ConsumerProducerDemo() {

	var channels [10]chan<- int

	// consumer
	for i := 0; i < 10; i++ {
		channels[i] = GorutineWorker(i)
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
	ConsumerProducerDemo()
}
