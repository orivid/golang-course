package main

import (
	"fmt"
	"math"
)

func main() {

	var smallInt int8 = math.MaxInt8
	var smallUInt uint = math.MaxUint8

	fmt.Printf("The largest number of smallInt is: %v\n", smallInt)
	fmt.Printf("The largest number of smallUint is: %v\n", smallUInt)

}
