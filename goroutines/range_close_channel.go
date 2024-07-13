package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channel")
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("Sent ", i)
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	// for {
	// 	time.Sleep(time.Second)
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
}
