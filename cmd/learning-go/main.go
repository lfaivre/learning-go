package main

import (
	"fmt"
	"reflect"
)

// utilizes compositional relationships, "has a" vs inheritance "is a"
type Animal struct {
	Height int
	Width  int
	Weight float32
}

// utilizing compositional embeedding
type Dog struct {
	Animal
	Name   string `required max:"100"`
	Age    int
	Breed  string
	Family []string
}

func main() {
	aDog := Dog{
		Name:  "Barker",
		Age:   1,
		Breed: "Golden Retriever",
		Family: []string{
			"Buster",
			"Frodo",
		},
	}
	fmt.Println(aDog)
	fmt.Println(aDog.Name)
	fmt.Println(aDog.Family[1])

	// utilizing composition
	aDog.Height = 32
	aDog.Width = 64
	aDog.Weight = 100
	fmt.Println(aDog)

	// anonymous struct
	aCat := struct{ name string }{name: "Whiskers"}
	fmt.Println(aCat)
	fmt.Println(aCat.name)

	// structs use values, not references
	anotherCat := aCat
	anotherCat.name = "Socks"
	fmt.Println(aCat.name)
	fmt.Println(anotherCat.name)

	// working with tags
	t := reflect.TypeOf(Dog{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	// tags are meaningless by themselves (just strings)
	// needs a validation framework to handle tag values
}
