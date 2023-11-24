package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 独立的消息生成器（服务、任务、组件）
func msgGen() <-chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message %d", i)
			i++
		}
	}()

	return c
}

func main() {
	// 类似于 c 语言中服务的句柄 handle
	m1 := msgGen()
	m2 := msgGen()
	for i := 0; i < 100; i++ {
		fmt.Printf("m1 = %v \n", <-m1)
		fmt.Printf("m2 = %v \n", <-m2)
	}
}
