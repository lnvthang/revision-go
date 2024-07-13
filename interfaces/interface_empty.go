package main

import "fmt"

// %v: get value, %T: get type
func main() {
	var i interface{}
	describe(i)

	i = "Go!"
	describe(i)

	i = 19
	describe(i)

	// Check type
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
