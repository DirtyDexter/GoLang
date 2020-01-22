package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x(2))
}

func foo() func(a int) int {
	return func(a int) int {
		return 12
	}
}
