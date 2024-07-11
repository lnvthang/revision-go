package main

import "fmt"

// Brief
// Pointer: Reference type and Value type
// &: indicate memory position (HEX)
// *: get value of pointer (value)

func main() {
	a, b := 77, 100

	// Pointer: Reference type
	p := &a
	// fmt.Println(p)
	// fmt.Println(*p)

	p = &b
	*p = 999
	// fmt.Println(p)
	// fmt.Println(*p)

	// Pointer: Value type
	z := b  // z := &b
	z = 500 // *z = 500
	fmt.Println(&z, &b)

}
