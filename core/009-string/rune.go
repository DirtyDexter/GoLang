package main

import "fmt"

func main() {
	x := "Hello"
	//fmt.Println(x)

	xRune := []rune(x)
	for i := 0; i < len(x); i++ {
		fmt.Println(string(xRune[i]))
	}
}
