package main

import (
	"fmt"
	"math/rand"
)

var randomNum int

func init() {
	fmt.Println("This is the second program initialization")
}

func init() {
	fmt.Println("This is where my program is initialized")
	randomNum = rand.Intn(251)
}

func main() {

	fmt.Printf("The value of randomNum is %v\n", randomNum)

	switch {
	case randomNum <= 100:
		fmt.Println("Between 0 and 100")
	case randomNum > 100 && randomNum <= 200:
		fmt.Println("Between 101 and 200")
	default:
		fmt.Println("Between 201 and 250")
	}

	for i := 0; i < 42; i++ {
		num := rand.Intn(5)

		switch num {
		case 0:
			fmt.Printf("The num is 0 for the %v time\n", i)
		case 1:
			fmt.Printf("The num is 1 for the %v time\n", i)
		case 2:
			fmt.Printf("The num is 2 for the %v time\n", i)
		case 3:
			fmt.Printf("The num is 3 for the %v time\n", i)
		case 4:
			fmt.Printf("The num is 4 for the %v time\n", i)
		}
	}

}
