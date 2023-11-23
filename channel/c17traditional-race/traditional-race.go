package main

import (
	"fmt"
	"time"
)

// 创建 int 的别名
type AtomicInt int

func (a *AtomicInt) increament() {
	*a++
}

func (a *AtomicInt) get() int {
	return int(*a)
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
