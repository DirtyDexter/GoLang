package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	last  string
	age   int
}

type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	x := []int{4, 1, 6, 9, 3}
	sort.Ints(x)
	fmt.Println(x)

	y := []string{"rk", "as", "pa", "ed"}
	sort.Strings(y)
	fmt.Println(y)

	people := []person{
		{"miss", "moneypenny", 29},
		{"rk", "ag", 27},
		{"james", "bond", 32},
	}
	fmt.Println(people)

	sort.Sort(byAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age > people[j].age
	})

}
