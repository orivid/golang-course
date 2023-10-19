package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	EB
)

const (
	FIRST = 2 + 2 + iota
	SECOND
	THIRD  = 8 + 8
	FOURTH = THIRD + 2
	FIFTH
)

func main() {

	fmt.Printf("size of KB in decimal: %v\t\tin binary: %b\n", KB, KB)
	fmt.Printf("size of MB in decimal: %v\t\tin binary: %b\n", MB, MB)
	fmt.Printf("size of GB in decimal: %v\tin binary: %b\n", GB, GB)
	fmt.Printf("size of TB in decimal: %v\tin binary: %b\n", TB, TB)
	fmt.Printf("size of EB in decimal: %v\tin binary: %b\n", EB, EB)

}
