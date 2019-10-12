package handler

import (
	"class-golang/golangforlife/Introduction/model"
	"fmt"
)

func Create() *model.Person {

	person := &model.Person{}
	dream := &model.Dream{}

	dream = &model.Dream{
		Type:        "wet",
		Description: "banana",
	}

	person.ID = 1
	person.Address = "jkt"
	person.Age = 25
	person.Name = "yana"
	person.Dream = *dream

	fmt.Printf(" \n\n\n model created %v", *person)

	return person

}

func Edit(person *model.Person) {
	fmt.Printf("\n\n before edit %v \n\n", *person)

	personNew := &model.Person{}

	personNew = &model.Person{
		ID:      person.ID,
		Name:    "myn",
		Age:     1,
		Address: "bks",
		Dream:   person.Dream,
	}

	fmt.Printf("\n\n after edit %v \n\n", *personNew)

}
