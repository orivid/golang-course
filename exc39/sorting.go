package main

import (
	"fmt"
	"slices"
)

func main() {

	nums := []int{7, 5, 10, 1, 2, 5, 20}

	fmt.Printf("Unsorted slice %v\n", nums)
	slices.Sort(nums)
	fmt.Printf("Sorted slice %v\n", nums)

}
