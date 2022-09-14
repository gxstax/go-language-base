package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	worker(c)
}

func worker(c chan int) {
	select {
	case <-c:
	// ... do some stuff
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
		return
	}
}
