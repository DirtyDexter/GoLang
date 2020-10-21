package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	var y *int = &x
	fmt.Println(y)

	foo(&x)
	fmt.Println(x)
}

func foo(y *int) {
	*y = 200
}
