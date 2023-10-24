package main

import "fmt"

func main() {

	fmt.Println(foo())
	num, str := bar()

	fmt.Println(num, str)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "My name is ori"
}
