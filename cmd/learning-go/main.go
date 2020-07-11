package main

import (
	"fmt"
)

func main() {
	// slice
	a := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// slices are reference types
	b := a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // slice of all elements
	e := c[3:]  // slice from 4th element to end (first number is inclusive)
	f := c[:6]  // slice first 6 elements (second number is exclusive)
	g := c[3:6] // slice from 4th, 5th, and 6th elements
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	h = append(h, 1, 2, 3, 4, 5)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	h = append(h, []int{6, 7, 8}...) // similar to javascript's spread operator
	fmt.Println(h)

	// treating slices like stacks
	i := []int{1, 2, 3, 4, 5}
	j := i[:len(i)-1] // remove last element
	fmt.Println(i)
	fmt.Println(j)

	k := append(i[:2], i[3:]...) // remove middle element
	fmt.Println(k)

	// warning: k has now modified i
	fmt.Println(i)
}
