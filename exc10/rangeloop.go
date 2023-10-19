package main

import (
	"fmt"
)

func main() {

	nums := []int{42, 43, 44, 45, 46}

	for i, n := range nums {
		fmt.Printf("The index is: %v and the value is: %v\n", i, n)
	}

	fmt.Printf("The value at index 2 is %v\n", nums[2])

	nameToAge := map[string]int{
		"ori":    38,
		"nuphar": 31,
	}

	for name, age := range nameToAge {
		fmt.Printf("The name is: %v and the age is: %v\n", name, age)
	}

	fmt.Printf("The age of ORI is: %v\n", nameToAge["ORI"])
	fmt.Printf("The age of ori is: %v\n", nameToAge["ori"])

	if age, ok := nameToAge["ori"]; ok {
		fmt.Printf("The age of ori is: %v when using comma ok\n", age)
	} else {
		fmt.Println("ori key was not found")
	}

}
