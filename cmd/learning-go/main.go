package main

import (
	"fmt"
)

func main() {
	var a int = 42
	// b is a pointer to an integer, it is assigned the address value of a (which is an integer)
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Printf("%v, %T\n", b, b)

	// dereferencing pointer to get integer value
	fmt.Printf("%v, %T\n", *b, *b)

	// use dereferenced value to change original
	*b = 36
	fmt.Println(a, *b)

	// pointer arithmetic is not native to go to keep the language simple
	// "unsafe" package allows for operations like pointer arithmetic

	type myStruct struct {
		foo int
	}

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// no initialization with `new` keyword
	var ms2 *myStruct
	ms2 = new(myStruct)
	fmt.Println(ms2)

	// `*` has lower precendence than `.` so parentheses are necessary
	(*ms2).foo = 42
	fmt.Println((*ms2).foo)

	// go compiler will implicity use value of struct, not pointer (shorthand)
	fmt.Println(ms2.foo)

	// zero value for a pointer is `nil`

	// all assignment operations in go are copy operations
	// slices and maps both pass values by reference
}
