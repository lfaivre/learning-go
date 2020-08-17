package main

import (
	"fmt"
)

/*

Overview

- Basic Syntax
- Parameters
- Return Values
- Anonymous Functions
- Functions as Types (functions are first-class citizens)
- Methods

*/

func main() {
	sayMessage("hello world!")

	// working with pointers
	greeting := "Hello"
	name := "Charlie"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	// working with variatic parameters
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Snoopy"
	fmt.Println(*name)
}

func sum(values ...int) *int {
	// values is a slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	// go promotes variable to shared/heap memory since it is being used outside of function's scope
	return &result
}
