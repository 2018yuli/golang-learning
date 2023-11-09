package main

import (
	"fmt"
	"time"
)

// chan 做为参数传递

func worker(c <-chan int) {
	// 收数据
	for {
		n := <-c
		fmt.Printf("n= %v \n", n)
	}
}

func channelDemo3() {
	// 创建 channel
	c := make(chan int)
	// 向 channel 中发数据
	go worker(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo3()
}
