package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("after error")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Custom error: Negative value")
	}
	return 12, nil
}
