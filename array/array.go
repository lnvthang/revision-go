package main

import "fmt"

var (
	arrStr []string
	arrNum []int
)

func main() {
	// fmt.Println(arrStr)
	arrNum := []int{1, 2, 3, 4, 5, 6}

	a := arrNum[0:2]
	b := arrNum[1:3]
	fmt.Println(a, b)

	b[0] = 99
	fmt.Println(a, b)
	fmt.Println(arrNum) // Ouput: [1 99 3 4 5 6]

	s1 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
	}

	s2 := s1[:2] // Output: s1[0:2]
	s3 := s1[1:] // Output: s1[1:the last position]

	fmt.Println(s1, s2, s3)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

}
