package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	languages map[string]string
}

func main() {
	per := person{firstName: "ori", lastName: "vider", languages: map[string]string{"hebrew": "heb"}}
	fmt.Println(per)
	rename("dan", per)
	fmt.Println(per)
}

func rename(firstName string, per person) {
	per.firstName = firstName
	per.languages["english"] = "eng"
}
