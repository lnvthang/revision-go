package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func myFunc2() (x, y int) {
	x = 9
	y = 10 - x
	return
}

// bool
// string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64
// byte = uint8
// rune = int32
// float32, float64
// complex32, complex64
// default value: 0, false, ""

var (
	s    = "Hello"
	n    = 1
	bool = true
)

const (
	M1 = "Mazda"
	M2 = "Mercedes"
	T1 = "Toyota"
)

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(myFunc2())

	fmt.Println(s, n, bool)

	x, y, z := false, "Go", 19

	fmt.Println(x, y, z)

	fmt.Println(M1, M2, T1)
}
