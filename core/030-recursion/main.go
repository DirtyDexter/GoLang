package main

import "fmt"

func main() {
	x := fact(5)
	fmt.Println(x)
}

func fact(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	return x * fact(x-1)
}
