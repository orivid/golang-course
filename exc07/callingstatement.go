package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x and y are %v %v", x, y)
	switch {
	case x < 4 && y < 4:
		fmt.Printf("x and y are smaller than 4")
	case y != 5:
		fmt.Printf("y is not equal to 5")
	default:
		fmt.Printf("Non of the conditions were met")
	}
}
