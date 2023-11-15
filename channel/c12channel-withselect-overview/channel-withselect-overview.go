package main

import "fmt"

func Generate() {

}

func SelectOverview() {
	// select 的使用场景，我们有多个 channel
	var c1, c2 chan int
	// 当我们想从多个 channel 中收取数据，谁快用谁，这时，可以使用select
	for {
		select {
		case n := <-c1:
			fmt.Printf("\n received value from c1 %d \n", n)
		case n2 := <-c2:
			fmt.Printf("\n received value from c2 %d \n", n2)
		default:
			fmt.Print(" ... ")
		}
	}

}

func main() {
	SelectOverview()
}
