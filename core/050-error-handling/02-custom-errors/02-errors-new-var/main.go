package main

import (
	"errors"
	"fmt"
	"log"
)

var errorCustom = errors.New("custom error: square root of negative number")

func main() {
	fmt.Printf("%T\n", errorCustom)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errorCustom
	}
	return 42, nil
}
