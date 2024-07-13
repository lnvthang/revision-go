package main

import (
	"fmt"
	"time"
)

// cap(), len()
func main() {
	fmt.Println("Channel")

	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("Sent ", i)
		}
	}()

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
		time.Sleep(time.Second)
	}
}
