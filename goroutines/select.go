package main

import (
	"fmt"
	"time"
)

// Select: the channel is assigned value that will print first
// return to out case and for

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	for i := 1; i <= 2; i++ {
		select {
		case c1 := <-ch1:
			fmt.Println("c1", c1)
		case c2 := <-ch2:
			fmt.Println("c2", c2)
		}
	}

}
