package handler

import (
	"class-golang/golangforlife/Introduction/model"
	"fmt"
)

func Create() *model.Person {

	person := &model.Person{}

	dream := &model.Dream{
		Type:        "wet",
		Description: "banana",
	}

	person.ID = 1
	person.Address = "jkt"
	person.Age = 25
	person.Name = "yana"
	person.Dream = *dream

	fmt.Printf("model created %v\n", *person)

	return person

}

func Edit(person *model.Person) {
	fmt.Printf("before edit %v", *person)

	personNew := &model.Person{}

	personNew = &model.Person{
		ID:      person.ID,
		Name:    "myn",
		Age:     1,
		Address: "bks",
		Dream:   person.Dream,
	}

	fmt.Printf("\nafter edit %v\n", *personNew)

}
