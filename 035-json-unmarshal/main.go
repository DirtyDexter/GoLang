package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"LastName"`
	Last  string `json:"FirstName"`
}

func main() {
	s := `[{"FirstName":"James","LastName":"Bond"}]`
	bs := []byte(s)

	var p1 []person

	err := json.Unmarshal(bs, &p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)
}
