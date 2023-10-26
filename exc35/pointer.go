package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	num := 42
	fmt.Println(&num)
	fmt.Println(num)

	p := "Drop by drop"
	q := "Persistently, patently"
	r := "The meaning of life is"
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {

	*a = "5"
	fmt.Println(*a, *b, *c, *d)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
