package main

import "fmt"

type Cars struct {
	Name  string
	Color string
	Price float32
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	// fmt.Println(Cars{"Mercedes", "Black", 123.78})

	// mercedes := Cars{"Mercedes", "Black", 123.78}
	// mercedes.Name = "Roll Royce"
	// fmt.Println(mercedes)

	// mazda := &mercedes
	// mazda.Color = "White"
	// fmt.Println(mercedes)

	fmt.Println(v1, p, v2, v3)
}
