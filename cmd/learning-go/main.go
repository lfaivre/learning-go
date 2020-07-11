package main

import (
	"fmt"
)

func main() {
	// map example
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      39250017,
		"Florida":    39250017,
		"New York":   39250017,
		"Illinois":   39250017,
		"Ohio":       39250017,
	}
	fmt.Println(statePopulations)

	// map example using make
	statePopulations2 := make(map[string]int)
	fmt.Println(statePopulations2)

	// get length of map
	fmt.Println(len(statePopulations))

	// access value of map element using key
	fmt.Println(statePopulations["Ohio"])

	// add key/value to map
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	// delete key/value from map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// ensure key/value exists using ok
	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)

	// maps pass value by reference
	sp := statePopulations
	fmt.Println(statePopulations)
	fmt.Println(sp)
	delete(sp, "Ohio")
	fmt.Println(statePopulations)
	fmt.Println(sp)

}
