package main

import (
	"fmt"
	"time"
)

// fatal error: all goroutines are asleep - deadlock!
func ChannelDeadLockNoGoroutine() {
	// var c chan int c == nil
	// 如何创建 chan
	c := make(chan int)

	// 向 channel 中发数据
	c <- 1
	c <- 2
	// 收数据
	n := <-c
	fmt.Printf("n= %v \n", n)

}

// fatal error: all goroutines are asleep - deadlock!
func ChannelDeadLockGoroutineExit() {
	// 创建 channel
	c := make(chan int)
	go func() {
		// 当 c 收完 n=1 之后
		// 此 goroutine 退出，造成 deadlock
		n := <-c
		fmt.Printf("n= %v \n", n)
	}()

	// 向 channel 中发数据
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}

// channel 是 goroutinue 和 goroutinue 之间通信管道
func ChannelIsGoroutine2Goroutine() {
	// 创建 channel
	c := make(chan int)
	go func() {
		// 收数据
		for {
			n := <-c
			fmt.Printf("n= %v \n", n)
		}
	}()

	// 向 channel 中发数据
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}

func main() {
	ChannelIsGoroutine2Goroutine()
}
