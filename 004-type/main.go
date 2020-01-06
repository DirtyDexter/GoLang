package main

import "fmt"

func main() {
	x := 3
	fmt.Printf("%T\n", x)
	var y = "Hello!!"
	fmt.Printf("%T\n", y)
	z := `Hi, "Steve"
	How r u?`
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	k := 1.23
	fmt.Println(k)
	fmt.Printf("%T\n", k)
}
