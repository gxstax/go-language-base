package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	println("woker is working")
	time.Sleep(2 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()

	return c
}

// 无缓存 channel 通信用作信号传递
func main() {
	println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
