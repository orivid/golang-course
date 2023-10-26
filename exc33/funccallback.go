package main

import (
	"fmt"
	"strconv"
)

func main() {

	isPositive(3, printHello)
	isPositive(-1, printHello)
}

func isPositive(num int, f func(s string)) {

	if num >= 0 {
		f(strconv.Itoa(num))
	}

}

func printHello(s string) {
	fmt.Println("The number", s, "is positive")
}
