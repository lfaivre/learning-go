package main

import (
	"fmt"
)

/*
primitive: boolean type

	values are true/false
	not an alias for other types
	zero value is false

	primitive: numeric types

integers
	signed integers
		int type has varying size, but min 32 bits
		8 bit (int8) through 64 bit (int64)
	unsigned integers
		8 bit (byte and uint8) through 32 bit (uint32)
	arithmetic operators
		addition, subtraction, multiplication, division, remainder
	bitwise operators
		and, or, xor, and not
	zero value is 0
	can't mix types in the same family

	int8: -128 - 127
	uint8: 0 - 255
	int16: -32,768 - 32,767
	uint16: 0 - 65,535
	int32: -2,147,483,648 - 2,147,483,647
	uint32: 0 - 4,294,967,295
	int64: -9,223,372,036,854,775,808 - 9,223,372,036,854,775,807

	AND: &
	OR: |
	XOR: ^
	ANDNOT: &^

	Bit Shift: << and >>

floating point numbers
	follow ieee-754 standard
	zero value is 0
	32 and 64 bit versions
	literal styles
		decimals (3.14)
		exponential (13e18 or 2e10)
		mixed (13.7e12)
	arithmetic operations
		addition, subtraction, multiplication, division

	float32: +-1.18E-38 - +-3.4E38
	float64: +-2.23E-308 - +-1.8E308

complex numbers
	zero value is (0+0i)
	64 and 128 bit versions
	Built-in functions
		complex - make complex number from two floats
		real - get real part as float
		imag - get imaginary part as float

primitive: strings
	strings
		utf-8
		immutable
		can be concatenated with plus (+) operator
		can be converted to []byte
	rune
		utf-32
		alias for int32
		special methods normally required to process
			strings.Reader#ReadRune
*/

// go does not do implicit operations, need to convert

func main() {
	// initializing defaults to zero value (boolean: false)
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	// complex numbers
	var m complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", real(m), real(m))
	fmt.Printf("%v, %T\n", imag(m), imag(m))

	var l complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", real(l), real(l))
	fmt.Printf("%v, %T\n", imag(l), imag(l))

	var k complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", real(k), real(k))
	fmt.Printf("%v, %T\n", imag(k), imag(k))

	// strings (UTF-8)
	// - aliases for bytes
	s := "Hello World"
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// slice of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// rune
	// - type alias for int32
	var r rune = 'h'
	fmt.Printf("%v, %T\n", r, r)
}
