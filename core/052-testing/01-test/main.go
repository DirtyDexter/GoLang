package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4, 5))
}

func sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
