package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("Hi my name is %v %v and I am %v\n", p.lastName, p.firstName, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		firstName: "ori",
		lastName:  "vider",
		age:       10,
	}

	//saySomething(p)
	saySomething(&p)

}
