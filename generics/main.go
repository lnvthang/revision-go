package main

import "fmt"

type Number interface {
	int64 | float64
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Sum[K comparable, V Number](m map[K]V) float64 {
	var s float64
	// Iterate over the map
	for _, value := range m {
		s += float64(value)
	}
	return s
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	map1 := map[string]int64{
		"M1": 30,
		"M2": 90,
	}

	map2 := map[string]int64{
		"M3": 50,
		"M4": 60,
	}
	fmt.Println(Sum(map1))
	fmt.Println(Sum(map2))
}
