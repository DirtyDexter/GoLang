package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {
	p1 := person{
		"James",
		"Bond",
	}

	people := []person{p1}

	pj, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)
	fmt.Println(string(pj))
}
