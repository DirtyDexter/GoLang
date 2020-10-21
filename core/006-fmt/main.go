package main

import "fmt"

func main() {
	var x = 1103
	fmt.Println(x)
	fmt.Printf("%v\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%o\n", x)

	fmt.Printf("%b\t%x\t%o\n", x, x, x)

	s := fmt.Sprintf("%b\t%x\t%o", x, x, x)
	fmt.Println(s)
}
