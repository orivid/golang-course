package main

import (
	"fmt"
	"github.com/orivid/golang-animals/dog"
)

func main() {
	barking := dog.Bark()
	barks := dog.Barks()
	fmt.Println(barking)
	fmt.Println(barks)

	fmt.Println(dog.BigBark())
}
