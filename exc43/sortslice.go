package main

import (
	"fmt"
	"slices"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	slices.Sort(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	slices.Sort(xs)
	fmt.Println(xs)
	slices.Reverse(xs)
	// sort xs
	fmt.Println(xs)
}
