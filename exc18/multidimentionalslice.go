package main

import (
	"fmt"
)

func main() {
	jb := [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."}}

	for i, value := range jb {

		fmt.Printf("The slice is %#v\n", value)
		for j, val := range value {
			fmt.Printf("The value of [%v,%v] is %v\n", i, j, val)
		}
	}
}
