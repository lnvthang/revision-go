package main

import "fmt"

func main() {
	printLine := func(msg string) string {
		return msg
	}

	fmt.Println(printLine("Hello, Golang!"))
}
