package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10000

	worker := 5

	wg := new(sync.WaitGroup)

	queueCh := make(chan int, n)

	for i := 0; i < n; i++ {
		queueCh <- i
	}

	close(queueCh)

	for i := 1; i <= worker; i++ {
		wg.Add(worker)
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 20)
				fmt.Printf("Worker %d is crawling web url %d\n", count, v)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
