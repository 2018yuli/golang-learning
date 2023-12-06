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

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		
	}()

	return out
}

func main() {
	// 以下非阻塞
	p := InMemorySort(ArraySource(3, 2, 6, 7, 4, 5, 1))
	// 阻塞
	for v := range p {
		fmt.Println(v)
	}
}
