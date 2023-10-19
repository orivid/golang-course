package main

import "fmt"

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	c3 = iota // c0 == 0
	c4
	c5
	c6
	c7
	c8
	c9
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", c1, c1)
	fmt.Printf("%d \t %b\n", c1<<c1, c1<<c1)
	fmt.Printf("%d \t %b\n", c4<<c2, c4<<c2)
	fmt.Printf("%d \t %b\n", c4<<c6, c4<<c6)
	fmt.Printf("%d \t %b\n", c4<<c7, c4<<c7)
	fmt.Printf("%d \t %b\n", c4<<c8, c4<<c8)
	fmt.Printf("%d \t %b\n", c4<<c9, c4<<c9)
}
