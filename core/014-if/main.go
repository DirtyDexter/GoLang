package main

import "fmt"

func main() {
	x := 2
	if 2 == 4 {
		fmt.Println("001")
	} else if 3 == 4 {
		fmt.Println("002")
	} else if x == 2 {
		fmt.Println("003")
	}

	if y := 3; y == 3 { // declaration in if loop
		fmt.Println("004")
	}
}
