package main

import "fmt"

func main() {
	// block scope
	x := 2
	fmt.Println(x)
	fmt.Println(y) // global y
	var y = 3      // local y
	y = 4
	fmt.Println(y)
	fmt.Println(z)
	z := 8 // local scope z
	z = 6
	fmt.Println(z)
	foo()
	fmt.Println(z)
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i) // scope of i is inside for loop
	fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	z = 7 // outer z
	//y = 10 // scopr of y is inside main function
}

// whole file scope
var z = 5

// assign default value depending on type
// false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var a int

// assigning some value
var b string = "RK"

// global y
var y = 11
