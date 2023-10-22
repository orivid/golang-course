package main

import "fmt"

func main() {

	personToFavoriteThings := map[string][]string{
		"bond_james":  {`shaken, not stirred`, `martins`, `fast cars`},
		"money_penny": {`intelligence`, `literature`, `computer science`},
		"no_dr":       {`cats`, `ice cream`, `sunsets`},
	}

	for k, v := range personToFavoriteThings {
		for i, thing := range v {
			fmt.Printf("%v likes %v as #%v\n", k, thing, i)
		}
	}

	personToFavoriteThings["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}

	fmt.Println(personToFavoriteThings)

	delete(personToFavoriteThings, "no_dr")
	delete(personToFavoriteThings, "ori")
	fmt.Println(personToFavoriteThings)

}
