package main

import (
	"fmt"
	"slices"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ByAge []user
type ByLastName []user

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	slices.SortFunc(users, func(a user, b user) int {

		if a.Age == b.Age {
			return 0
		}

		if a.Age < b.Age {
			return -1
		} else {
			return 1
		}
	})

	slices.SortFunc(users, compare)

	fmt.Println("I am doing something")

	fmt.Println(users)

}

func compare(a user, b user) int {

	if a.Age == b.Age {
		return 0
	}

	if a.Age < b.Age {
		return -1
	} else {
		return 1
	}
}

func (b ByLastName) Len() int {
	return len(b)
}

func (b ByLastName) Less(i, j int) bool {
	return b[i].Last < b[j].Last
}

func (b ByLastName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[i].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
