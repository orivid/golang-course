package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func main() {
	var bookPtr *Book = &Book{pages: 5}
	// Call the two implicit declared functions.
	fmt.Println(*bookPtr)
	bookPtr.SetPages(123)

	fmt.Println(*bookPtr)
	fmt.Println(bookPtr.Pages()) // 123
}
