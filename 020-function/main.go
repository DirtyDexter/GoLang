package main

import "fmt"

func main() {
	result, numbers := sum(2, 7)
	fmt.Println(result, numbers)
}

func sum(x int, y int) (int, string) {
	return x + y, "no error"
}
