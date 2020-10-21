package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 5
	fmt.Println(x)

	var y [2][3]int
	fmt.Println(y)
	y[1][2] = 12
	fmt.Println(y)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
