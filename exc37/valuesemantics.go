package main

import "fmt"

type person struct {
	first string
}

type greeter struct {
	prefix string
}

func createGreeter() *greeter {
	return &greeter{"Hi"}
}

func (g greeter) greet() {
	fmt.Println(g.prefix, "my name is ori")
}

func main() {

	p1 := person{"ori"}
	p2 := person{"dan"}
	changePersonName("oriori", p1)
	changePersonName1("dandan", &p2)
	fmt.Println(p1)
	fmt.Println(p2)

	g := createGreeter()
	fmt.Println(g)
	g.greet()
}

func changePersonName(name string, p person) person {
	p.first = name
	return p
}

func changePersonName1(name string, p *person) {
	p.first = name
}
