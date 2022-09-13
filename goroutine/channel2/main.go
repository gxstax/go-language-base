package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// 无缓存 channel 类型的发送与接收
func main() {
	//withBuffer()
	sendOrRecvOnly()
}

// 不带缓冲区的 channel
func withoutBuffer() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 13
	}()

	go func() {
		t := <-ch1
		logrus.Println(t)
	}()

	n := <-ch1

	logrus.Println(n)
}

// 带缓冲区的 channel
func withBuffer() {
	ch2 := make(chan int, 1)

	ch2 <- 10
	n := <-ch2
	fmt.Println(n)

	ch3 := make(chan int, 2)
	ch3 <- 17
	ch3 <- 18
}

// 只发送/只接收 channel
func sendOrRecvOnly() {
	// 只发送 channel
	ch1 := make(chan<- int, 1)
	// 只接收 channel
	ch2 := make(<-chan int, 1)

	// 只发送 channel 接收报错
	//<-ch1	// Invalid operation: <-ch1 (receive from the send-only type chan<- int)

	// 只接收 channel 发送报错
	//ch2 <- 10 // Invalid operation: ch2 <- 10 (send to the receive-only type <-chan int)

	ch1 <- 10
	fmt.Println(ch1)

	x := <-ch2
	fmt.Println(x)
}
