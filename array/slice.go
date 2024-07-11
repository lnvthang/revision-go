package main

import "fmt"

// len(), cap()
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	a_slice := make([]int, 0, 5)
	fmt.Println(a_slice, len(a_slice), cap(a_slice))

	b_slice := append(a_slice, 0) // Return the new slice
	fmt.Println(a_slice, b_slice)
}
