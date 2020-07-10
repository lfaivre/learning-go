package main

import (
	"fmt"
)

// constants are immutable, but can be shadowed

// constants must be calculable at compile-time
// - can't assign a function that returns a value to a constant

// typed constants work like immutable variables
// can utilize implicit conversions for untyped constants (with similar types)

// enumerated constant
const a = iota // starts at 0, increments by one

const (
	_ = iota // ignore first value by using blank identifier
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	// defining a constant
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// enumerated constant example
	var specialistType int = dogSpecialist
	fmt.Printf("%v\n", specialistType == dogSpecialist)
	fmt.Printf("%v\n", specialistType == catSpecialist)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
