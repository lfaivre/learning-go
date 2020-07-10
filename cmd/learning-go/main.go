package main

import (
	"fmt"
	"strconv"
)

// naming conventions:
// - lowercase variable names are scoped to the package
// - uppercase variable names are exposed to outside
// - use verbose variable names for variables that are used outside of package scope

// use full syntax when declaring variable at package-level
var l int = 42

// multiple declarations
var (
	string1 string = "this is string1"
	string2 string = "this is string2"
)

var s int = 0

func main() {
	// first way to initialize variable
	var i int
	i = 42

	// second way to initialize variable
	var j int = 27

	// third way to initialize variable
	k := 99

	// can't redeclare variables
	// shadowing, variables declared in inner-scope take precedence

	s := 32

	// type conversions
	var t int = 42
	var u float32
	u = float32(t)

	v := strconv.Itoa(t)

	// all local variables must be used
	fmt.Println(i, j, k, s, u, v)
}
