package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr {
		fmt.Println("index: ", i, ", item", v)
	}

	for i := range arr {
		fmt.Println("index: ", i)
	}

	for _, v := range arr {
		fmt.Println("item: ", v)
	}
}
