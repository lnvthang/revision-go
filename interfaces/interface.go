package main

import "fmt"

type MyInterface interface {
	PrintMsg() string
}

type Number interface {
	int64 | float64
}

func main() {
	myfunc1 := MyFunc1("MyFunc1 implemented MyInterface")
	fmt.Println(myfunc1.PrintMsg())

	myfunc2 := MyFunc1("MyFunc2 implemented MyInterface")
	fmt.Println(myfunc2.PrintMsg())

	var a MyInterface
	a = myfunc1
	a = myfunc2
}

type MyFunc1 string // Type non-structure, anonymous structure

// Same name and return value => Implemented interface
func (f MyFunc1) PrintMsg() string {
	return string(f)
}

type MyFunc2 struct {
	Msg2 string
}

func (s MyFunc2) PrintMsg() string {
	return s.Msg2
}
