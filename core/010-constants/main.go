package main

import "fmt"

const (
	a        = 32
	b string = "test"
)

const c = 23.233

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\t%T\t%T\n", a, b, c)
}
