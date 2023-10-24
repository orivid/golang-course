package main

import "fmt"

func main() {

	numbers := []int{1, 3, 5, 7, 9}
	sum := foo(numbers...)
	fmt.Println("The sum of the numbers is", sum)
	secondSum := bar(numbers)
	fmt.Println("The second sum of numbers is", secondSum)
}

func foo(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func bar(nums []int) int {
	return foo(nums...)
}
