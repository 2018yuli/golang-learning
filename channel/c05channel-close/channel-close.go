package main

import (
	"fmt"
	"time"
)

func GorutineWorkerReceiveZeroForever(i int, c <-chan int) {
	for {
		// 当 c 被 close 以后，接收方将不停的收到 “空” 值（int 对应 0）
		fmt.Printf("worker %d received data %c \n", i, <-c)
	}
}

func GorutineWorkerJudgeChannelClosed(i int, c <-chan int) {
	for {
		n, ok := <-c
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		// 当 c 被 close 以后，接收方将不停的收到 “空” 值（int 对应 0）
		fmt.Printf("worker %d received data %c \n", i, n)
	}
}

func GorutineWorkerStopWhenChannelClosed(i int, c <-chan int) {
	for n := range c {
		fmt.Printf("worker %d received data %c \n", i, n)
	}
}
// Communication Sequential Process
func channelCanbeClosed() {
	c := make(chan int)
	go GorutineWorkerStopWhenChannelClosed(3, c)
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// not all channels need to be closed!
	channelCanbeClosed()
}
