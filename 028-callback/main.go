package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenSum := even(sum, a...)
	fmt.Println(evenSum)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(sum func(x ...int) int, x ...int) int {
	var even []int
	for _, v := range x {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return sum(even...)
}
