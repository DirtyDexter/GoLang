package main

import "fmt"

func main() {
	defer foo()
	bar()

	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println(f())
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
