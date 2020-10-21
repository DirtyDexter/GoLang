package main

import "fmt"

type person struct {
	first string
	last  string
}

type agent struct {
	person
	ltk bool
}

type human interface {
	speak(greeting string)
}

func main() {
	p1 := person{
		"rahul",
		"kumar",
	}
	fmt.Println(p1)
	p1.speak("Hello")

	a1 := agent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	fmt.Println(a1)
	a1.speak("Hi")

	bar(a1)
	bar(p1)

	bar2(a1)
	bar2(p1)
}

func (a agent) speak(greeting string) {
	fmt.Println(greeting, "i am agent", a.first, a.last)
}

func (p person) speak(greeting string) {
	fmt.Println(greeting, "i am person", p.first, p.last)
}

func bar(h human) {
	fmt.Println("i am passed in bar - ", h)
}

func bar2(h human) {
	switch h.(type) {
	case person:
		fmt.Println("i am passed in bar2 and i m person", h.(person).first)
	case agent:
		fmt.Println("i am passed in bar2 and i m agent", h.(agent).first)
	}
}
