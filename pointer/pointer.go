package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // get memory position of i (reference)
	fmt.Println(p)  // read memory position i through the pointer (Hex)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	fmt.Println(j) // see the new value of j
}
