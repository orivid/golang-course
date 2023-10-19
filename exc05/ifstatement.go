package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNum := rand.Intn(251)

	fmt.Printf("The value of randomNum is %v\n", randomNum)

	if randomNum <= 100 {
		fmt.Println("Between 0 and 100")
	} else if randomNum > 100 && randomNum <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}
