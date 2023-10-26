package main

import "fmt"

func main() {

	addFunc := func(x int, y int) int {
		return x + y
	}

	fmt.Println(addFunc(3, 5))

	x := func(x int, y int) int {
		return x * y
	}(5, 6)

	fmt.Println(x)

	add := createAddFunc()
	fmt.Println(add(4, 6))
	fmt.Println(add(4, 6))

}

func createAddFunc() func(int, int) int {
	counter := 0
	return func(x int, y int) int {
		result := x + y + counter
		counter++
		return result
	}
}
