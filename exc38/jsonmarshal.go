package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	UserID    int    `json:"user_id"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	p := person{
		UserID:    111,
		ID:        555,
		Title:     "This is my title",
		Completed: false,
	}

	perBytes, err := json.Marshal(p)

	if err != nil {
		log.Fatalf("Error while trying to marshal person. %v", err)
	}

	fmt.Println(string(perBytes))

	var newPer person

	err = json.Unmarshal(perBytes, &newPer)

	if err != nil {
		_ = fmt.Errorf("there was an error while unmarshaling %v", err)
	}

	fmt.Println(newPer)

}
