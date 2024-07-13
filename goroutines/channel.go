package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channel")

	ch := make(chan int) // unbuffered

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
