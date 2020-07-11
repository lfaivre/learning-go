package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Student 1"
	students[1] = "Student 2"
	students[2] = "Student 3"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students Length: %v\n", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// arrays are passed by value, not reference
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// unless...
	c := [...]int{1, 2, 3}
	d := &c // address of c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
}
