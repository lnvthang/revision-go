package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(time.Second * 2)
	fmt.Println(s)
}

func main() {
	go say("world")
	say("hello")
}
