package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavour []string
}

func main() {

	firstPerson := person{
		firstName:       "ori",
		lastName:        "vider",
		iceCreamFlavour: []string{"mango"},
	}

	secondPerson := person{
		firstName:       "nuphar",
		lastName:        "pianko",
		iceCreamFlavour: []string{"chocolate"},
	}

	fmt.Println(firstPerson)
	fmt.Println(secondPerson)
	fmt.Println(firstPerson.getFirstName())

	lastNameToPerson := make(map[string]person)
	lastNameToPerson[firstPerson.lastName] = firstPerson
	lastNameToPerson[secondPerson.lastName] = secondPerson

	for lastName, p := range lastNameToPerson {
		for _, flavour := range p.iceCreamFlavour {
			fmt.Printf("%v likes %v\n", lastName, flavour)
		}
	}

}

func (person person) getFirstName() string {
	return person.firstName
}
