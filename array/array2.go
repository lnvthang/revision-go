package main

import "fmt"

func myFunc(num int16) int16 {
	if n := num; num > 0 {
		return n
	} else {
		return 0
	}
}

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

	sum := 1
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)

	// for {
	// 	break
	// }

	fmt.Println(myFunc(-1))
}
