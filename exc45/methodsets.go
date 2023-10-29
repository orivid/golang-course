package main

import (
	"fmt"
	"math"
)

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	var bookPtr *Book = &Book{pages: 5}
	// Call the two implicit declared functions.
	fmt.Println(*bookPtr)
	bookPtr.SetPages(123)

	fmt.Println(*bookPtr)
	fmt.Println(bookPtr.Pages()) // 123

	c := circle{radius: 2.5}
	info(&c)

}
