package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	age       int
}

type persons []person

func main() {

	p1 := person{
		firstName: "ori",
		age:       30,
	}

	p2 := person{
		firstName: "dan",
		age:       20,
	}

	per := []person{p1, p2}

	fmt.Println("Before sort", per)
	sort.Sort(persons(per))
	fmt.Println("After sort", per)

}

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p persons) Swap(i, j int) {
	tempI := p[i]
	tempJ := p[j]

	p[j] = tempI
	p[i] = tempJ
}
