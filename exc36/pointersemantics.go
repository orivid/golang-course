package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}
func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) // doesn't run

	d2 := dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)
	d1.Bark()
	youngRun(d1)
	d2 = d2.changeName("DDog")
	fmt.Println(d2)

}

func (d dog) Bark() {
	fmt.Println("My name is", d.first, "and I am barking")
}

func (d dog) changeName(firstName string) dog {
	d.first = firstName
	return d
}
