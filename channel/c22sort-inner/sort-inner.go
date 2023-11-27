package main

import (
	"fmt"
	"sort"
)

func SortInner() {
	a := []int{3, 6, 2, 1, 9, 8, 10, 4, 5, 7}
	sort.Ints(a)
	// fmt.Printf("a = %v", a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	SortInner()
}
