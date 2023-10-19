package main

import (
	"fmt"
)

func main() {
	var x int

	for x < 20 {

		fmt.Printf("The current number of x is: %v\n", x)
		x += 1
	}

	var y int

	for {

		if y >= 15 {
			break
		}

		y++
		fmt.Printf("The current number of y is: %v\n", y)
	}

}
