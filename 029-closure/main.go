package main

import "fmt"

func main() {
	add5 := add(5)
	fmt.Println(add5(3))
	fmt.Println(add5(5))
	fmt.Println(add5(6))
}

func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
