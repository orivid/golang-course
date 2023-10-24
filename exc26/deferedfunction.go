package main

import "fmt"

func main() {
	defer bar()
	defer fooBar()
	foo()
	fmt.Println("This is the end of the main func")
}

func foo() {
	fmt.Println("This is func foo calling")
}

func bar() {
	fmt.Println("This is func bar calling")
}

func fooBar() {
	fmt.Println("This is func foo bar calling")
}
