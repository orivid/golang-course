package main

import (
	"fmt"
)

var middleName = "meir"

func main() {

	const (
		_ = iota
		one
		two
	)

	var b bool
	var a string

	fmt.Printf("This is the value of a: %v\n", a)
	fmt.Printf("This is the value of b: %v\n", b)

	firstName := "ori"
	lastName := "vider"

	fmt.Printf(firstName + " " + lastName + "\n")

	city, country := "petach tikva", "israel"

	fmt.Printf("This is my city: %s \tand country: %s\n", city, country)
	fmt.Printf("This is my middle name: %s\n", middleName)
	fmt.Printf("1+2=%v", one+two)

}
