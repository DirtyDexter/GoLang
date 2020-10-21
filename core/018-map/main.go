package main

import "fmt"

func main() {
	x := map[string]string{"name": "rk"}
	x["city"] = "farah"
	fmt.Println(x)
	fmt.Println(x["name"])
	fmt.Println(x["address"])

	v, ok := x["city"]
	fmt.Println(v, ok)
	if v, ok := x["name"]; ok {
		fmt.Println(v)
	}

	for k, v := range x {
		fmt.Println(k, v)
	}

	y := make(map[string]string, 1)
	fmt.Println(y)
	fmt.Println(len(y))

	y["test"] = "test1"
	y["test2"] = "pass"
	fmt.Println(y)
	fmt.Println(len(y))

	delete(y, "test")
	fmt.Println(y)
	fmt.Println(len(y))

	delete(y, "test3")
	fmt.Println(y)
	fmt.Println(len(y))
}
