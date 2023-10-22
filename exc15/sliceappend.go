package main

import "fmt"

func main() {

	var nums []int
	num := 42

	for i := 0; i < 10; i++ {
		nums = append(nums, num)
		num++
	}

	nums = append(nums, 52)
	fmt.Println(nums)

	nums = append(nums, 53, 54, 55)
	fmt.Println(nums)
	newNums := append(nums, []int{56, 57, 58, 59}...)
	fmt.Println(newNums)

}
