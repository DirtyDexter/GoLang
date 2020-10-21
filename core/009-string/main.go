package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	x := "Hello"
	x = "Hello, world!!" // string is immutable
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%b\t", x[i])
	}

	fmt.Println("")

	for i := 0; i < len(x); i++ {
		fmt.Printf("%d\t", x[i])
	}

	fmt.Println("")

	for i := 0; i < len(x); i++ {
		fmt.Printf("%#U\t", x[i])
	}

	fmt.Println("")

	s := []byte(x) // string is sequence of bytes
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b\t", s[i])
	}

	fmt.Println("")

}
