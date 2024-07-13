package main

import (
	"fmt"
	"strings"
)

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func main() {
	num_arr := []int{1, 2, 3, 4, 5}

	rs := mapAny(num_arr, func(v int) int {
		return v * 2
	})

	str_arr := []string{"Go", "C"}

	rsStr := mapAny(str_arr, func(v string) string {
		return strings.ToUpper(v)
	})

	fmt.Println(rs)
	fmt.Println(rsStr)
}
