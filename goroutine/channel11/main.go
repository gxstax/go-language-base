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
	heartbeat := time.NewTicker(5 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case <-c:

		case <-heartbeat.C:
			fmt.Println("heartbeat C...")
		}
	}
}
