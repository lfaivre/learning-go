package main

import (
	"fmt"
	"math"
)

func returnTrue() bool {
	fmt.Println("Returning true!")
	return true
}

func main() {
	// basic if conditional statement
	if true {
		fmt.Println("The test is true!")
	}

	// conditional statement using initializer
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        39250017,
		"Florida":      20612439,
		"New York":     39250017,
		"Pennsylvania": 39250017,
		"Illinois":     39250017,
		"Ohio":         39250017,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// utilizing comparison operators
	number := 50
	guess := 20

	if guess < number {
		fmt.Println("Too low!")
	}

	// utilizing logical operators
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		fmt.Println("Executing second block!")
	}

	// utilizing short-circuiting, lazily evaluation as needed
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	// floating point numbers are an appoximation of decimal values
	myNum := 0.123

	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same!")
	} else {
		fmt.Println("These are different!")
	}

	// fine tune approximation value (0.001 in this case) in order to compare floating point nums
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same!")
	} else {
		fmt.Println("These are different!")
	}
}
