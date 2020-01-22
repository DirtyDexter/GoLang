package main

import "fmt"

func main() {
	x := func(a int) {
		fmt.Println("i am func expression", a)
	}
	x(2)
}
