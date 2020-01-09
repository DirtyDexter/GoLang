package main

import "fmt"

var x = 9223372036854775807

// var y int8 = 255 // will fail
var y int8 = -128

var z float64 = 9223372036854775807

var a int64 = 9223372036854775807

var b = 1233434.3444

var c uint16 = 65535

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
