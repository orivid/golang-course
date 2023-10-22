package main

import "fmt"

func main() {

	var nums []int
	num := 42

	for i := 0; i < 10; i++ {
		nums = append(nums, num)
		num++
	}

	nums = append(nums[:3], nums[6:]...)
	fmt.Println(nums, cap(nums), len(nums))

}
