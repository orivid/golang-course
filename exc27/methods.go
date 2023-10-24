package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func main() {
	pers := person{firstName: "Ori", age: 37}
	fmt.Println(pers.Speak())
}

func (p person) Speak() string {
	return fmt.Sprintf("Hi my name is %v and my age is %v", p.firstName, p.age)
}
