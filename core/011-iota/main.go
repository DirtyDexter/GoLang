package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c
		d
	)

	const (
		e = iota
		f = iota
	)

	const (
		g, h = iota, iota
		i, j
	)

	//var x = iota // error

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	fmt.Printf("%T\n", a)
}
