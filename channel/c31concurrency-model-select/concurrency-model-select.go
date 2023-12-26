package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 独立的消息生成器（服务、任务、组件）
func msgGen(name string) <-chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("%v generate message %d", name, i)
			i++
		}
	}()

	return c
}

func fanIn(chs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ch := range chs {
		// chCopy = ch
		go func(ch <-chan string) {
			for {
				c <- <-ch
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	// 类似于 c 语言中服务的句柄 handle
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m := fanIn(m1, m2)
	for i := 0; i < 100; i++ {
		fmt.Printf("m = %v \n", <-m)
	}
	m3 := msgGen("service1")
	m4 := msgGen("service2")
	m5 := fanInBySelect(m3, m4)
	for i := 0; i < 100; i++ {
		fmt.Printf("m5 = %v \n", <-m5)
	}
}
