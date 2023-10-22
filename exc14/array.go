package main

import "fmt"

func main() {

	nums := [5]int{}
	nums[0] = 1
	nums[1] = 4
	nums[4] = 2
	nums[3] = 6
	nums[2] = 7

	for _, val := range nums {
		fmt.Printf("The value is : %v the type is: %T\n", val, val)
	}

}
