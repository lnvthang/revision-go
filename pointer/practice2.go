package main

import "fmt"

func main() {
	a := 10
	z := &a
	fmt.Println(z)
	fmt.Println(*z)

	*z = 50
	fmt.Println(*z)
}
