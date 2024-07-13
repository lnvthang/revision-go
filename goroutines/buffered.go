package main

import (
	"fmt"
)

func main() {
	// FIFO
	ch := make(chan int, 1) // 1: size of channel
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)

	ch2 := make(chan int, 2)
	ch2 <- 3
	ch2 <- 4
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	ch3 := make(chan int, 2)
	ch3 <- 3
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
