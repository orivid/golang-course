package main

import "fmt"

func main() {

	person := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "ori",
		friends: map[string]int{
			"andrey": 40,
		},
		favDrinks: []string{
			"water",
		},
	}

	fmt.Println(person.friends)
	fmt.Println(person.favDrinks)

	for name, age := range person.friends {
		fmt.Printf("%v is %v years old\n", name, age)
	}
}
