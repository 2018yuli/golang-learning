package main

import (
	"fmt"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}

		close(out)
	}()

	return out
}

func main() {
	// 以下两步为非阻塞
	p := ArraySource(3, 2, 6, 7, 4, 5, 1)
	p = InMemorySort(p)
	// 阻塞
	for {
		if num, ok := <-p; ok {
			fmt.Printf("f=%d \r\n", num)
		} else {
			break
		}
	}
}
