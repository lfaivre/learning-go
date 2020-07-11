package main

import (
	"fmt"
)

func main() {
	// basic for-loop
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
	}

	// while loop (modified for-loop)
	for i := 0; i < 5; {
		// fmt.Println(i)
		i++
	}

	// while loop (modified for-loop)
	j := 0
	for {
		if j == 4 {
			break
		}
		j++
		fmt.Println(j)
	}

	// use labels to break out of multiple loops
Loop:
	for k := 1; k <= 3; k++ {
		for l := 1; l <= 3; l++ {
			fmt.Println(k * l)
			if k*l >= 3 {
				break Loop
			}
		}
	}

	// for-range loop
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value)
	}
}
