package main

import (
	"fmt"
	"sync"
	"time"
)

// Select: the channel is assigned value that will print first
// return to out case and for

func main() {
	lock := new(sync.Mutex)

	count := 0

	for i := 1; i <= 5; i++ {
		go func() {
			for y := 1; y <= 10; y++ {
				lock.Lock()
				count++
				fmt.Println(count)
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 5)
}
