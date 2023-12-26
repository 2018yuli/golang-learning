package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 独立的消息生成器（服务、任务、组件）done 用于服务中断
func msgGen(name string, done chan struct{}) <-chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("%v generate message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()

	return c
}

func nonBlockWait(c <-chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c <-chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	// 类似于 c 语言中服务的句柄 handle
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	m2 := msgGen("service2", done)
	for i := 0; i < 10; i++ {
		fmt.Printf("m1 = %v \n", <-m1)
		if m, ok := nonBlockWait(m2); ok {
			fmt.Printf("m2 = %v \n", m)
		} else {
			fmt.Printf("no message from service2\n")
		}
	}

	m3 := msgGen("service3", done)
	for i := 0; i < 10; i++ {
		if m, ok := timeoutWait(m3, 1*time.Second); ok {
			fmt.Printf("m3 = %v \n", m)
		} else {
			fmt.Printf("time out service3\n")
		}
	}
	done <- struct{}{}
	// time.Sleep(time.Second)
	// 优雅退出
	<-done
}
