package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		"james",
		"bond",
	}
	fmt.Println(p1)
	p1.speak("Hello")
}

func (p person) speak(greeting string) {
	fmt.Println(greeting, "i am", p.first, p.last)
}
