package main

import "fmt"

var a = 42

type hotdog int

var b hotdog

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	foo()
	b = hotdog(a) // conversion, not casting
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}

func foo() {
	type hotdog int
	var c hotdog = 23
	fmt.Printf("%T\n", c)
	fmt.Println(c)
}
