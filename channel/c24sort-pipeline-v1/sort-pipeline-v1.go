package main

import "fmt"

func ArraySource(a ...int) chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

func main() {
	p := ArraySource(3, 2, 6, 7, 4, 5, 1)
	for {
		if num, ok := <-p; ok {
			fmt.Printf("f=%d \r\n", num)
		} else {
			break
		}
	}
}
