package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5}
	sum := sum(x...)
	fmt.Println(sum)

	y := []int{1, 2, 3}
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	change(y...)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	z := make([]int, 3, 10)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	changeZ(z...)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}

func sum(x ...int) int {
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func change(x ...int) {
	fmt.Printf("%T\n", x) // underlying array wil lbe same as y
	x[0] = 4              // will change y also

	x = append(x, 5) // created new underlying array, no effect on y nowonwards
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 8 // will not chnage y
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func changeZ(x ...int) {
	fmt.Printf("%T\n", x) // underlying array wil lbe same as z
	x[0] = 4              // will change z also

	x = append(x, 5) // same underlying array becuase capacity is still <10
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 8 // will chnage z also
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
