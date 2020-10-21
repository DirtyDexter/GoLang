package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := append(x, 4)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	z := make([]int, 10, 10)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	xx := make([]int, 10)
	z = append(z, xx...)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	yy := [5]int{1, 2, 3, 4, 5}
	fmt.Println(yy[1:5])
	fmt.Println(yy[:3])
	fmt.Println(yy[3:])

	for i, v := range yy {
		fmt.Println(i, v)
	}

	zz := [2][3]int{{1, 2}, {1}}
	fmt.Println(zz)
	fmt.Println(len(zz))
	fmt.Println(cap(zz))
}
