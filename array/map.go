package main

import "fmt"

type TwoNums struct {
	X, Y int16
}

var (
	m map[string]TwoNums
)

var m2 = map[string]TwoNums{
	"Mazda": TwoNums{
		3, 4,
	},
	"Toyota": TwoNums{
		5, 6,
	},
}

func main() {
	m = make(map[string]TwoNums) // Dictionary: Key - Value
	m["Mercedes"] = TwoNums{
		1, 2,
	}

	fmt.Println(m["Mercedes"])
	fmt.Println(m)
	fmt.Println(m2["Mazda"])

	// Change value in map
	v, ok := m2["RollRoyce"]
	fmt.Println("The value: ", v, "Present?", ok)
}
