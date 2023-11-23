package main

import (
	"fmt"
	"sync"
	"time"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) increament() {

	// sync block
	fmt.Println("sync block")
	// 使用匿名函数实现代码块加锁
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()

}

func (a *AtomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func traditionalSync() {
	var a AtomicInt
	a.increament()
	go func() {
		a.increament()
	}()
	time.Sleep(time.Millisecond)
	fmt.Printf("a = %d \n", a.get())
}

// cd channel/c17traditional-race/
// go run -race traditional-race.go
// WARNING: DATA RACE：21 行写，15 行读
func main() {
	traditionalSync()
}
