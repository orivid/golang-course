package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  int
	model string
	doors int
	color string
}

func main() {

	firstVeh := vehicle{engine: engine{false},
		make:  1985,
		model: "challenger",
		doors: 4,
		color: "blue",
	}

	secondVeh := vehicle{engine: engine{true}, make: 2022, model: "Y", doors: 4, color: "white"}

	fmt.Printf("%#v\n", firstVeh)
	fmt.Printf("%#v\n", secondVeh)

	fmt.Printf("The make of the tesla is: %v\n", secondVeh.make)
	fmt.Printf("Is challenger an electric car: %v\n", firstVeh.electric)
}
