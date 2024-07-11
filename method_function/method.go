package main

import "fmt"

type NumberStruct struct {
	a, b int16
}

func SumNumber(num NumberStruct) int16 {
	return num.a + num.b
}

func main() {
	sum := NumberStruct{7, 9}
	fmt.Println(SumNumber(sum))
}
