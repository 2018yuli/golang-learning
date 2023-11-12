package main

import (
	"fmt"
	"time"
)

func GorutineWorker(i int, c <-chan int) {
	for {
		fmt.Printf("worker %d received data %c \n", i, <-c)
	}
}

func bufferedChannelForImprovePerformance() {
	// buffered[3]
	c := make(chan int, 3)
	go GorutineWorker(3, c)
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	time.Sleep(time.Millisecond)
}

func main() {
	bufferedChannelForImprovePerformance()
}
