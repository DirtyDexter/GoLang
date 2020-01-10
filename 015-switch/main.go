package main

import "fmt"

func main() {
	switch 2 + 3 {
	case 5:
		{
			fmt.Println("5")
		}
	case 6:
		{
			fmt.Println("6")
		}
	default:
		{
			fmt.Println("default")
		}
	}

	switch 3 + 3 {
	case 5, 6, 7:
		{
			fmt.Println("6")
		}
	default:
		{
			fmt.Println("default")
		}
	}

	switch {
	case false:
		{
			fmt.Println("false")
		}
	case true:
		{
			fmt.Println("true")
		}

	default:
		{
			fmt.Println("default")
		}
	}

	switch {
	case true:
		{
			fmt.Println("true1")
		}
	case true:
		{
			fmt.Println("true2")
		}

	default:
		{
			fmt.Println("default")
		}
	}

	switch {
	case true:
		{
			fmt.Println("true1")
		}
		fallthrough
	case false:
		{
			fmt.Println("true2")
		}
	default:
		{
			fmt.Println("default")
		}
	}

	switch {
	case true:
		{
			fmt.Println("true1")
		}
		fallthrough
	case false:
		{
			fmt.Println("true2")
		}
		fallthrough
	default:
		{
			fmt.Println("default")
		}
	}

	switch {
	case true:
		{
			fmt.Println("true11111")
		}
		fallthrough
	default:
		{
			fmt.Println("default11111")
		}
	case false:
		{
			fmt.Println("true2222222")
		}
	}
}
