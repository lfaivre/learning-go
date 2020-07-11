package main

import (
	"fmt"
)

func main() {
	// basic switch statement
	// "2" is referred to as a "tag"
	// "break" keyword is implied
	// can still use "break" however if needed
	switch 5 {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
		break
		fmt.Println("Will not run")
	default:
		fmt.Println("Default")
	}

	// switch statement with initializer
	switch i := 2 + 3; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
	default:
		fmt.Println("Default")
	}

	// tagless syntax
	// overlapping is allowed, first to equal true will evaluate
	// use "fallthrough" to fall through to next case
	j := 10
	switch {
	case j <= 10:
		fmt.Println("First")
	case j <= 20:
		fmt.Println("Second")
		fallthrough
	case j >= 0:
		fmt.Println("Greater than or equal to 0")
	default:
		fmt.Println("Default")
	}

	// type switch
	var k interface{} = 1.0
	switch k.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}
