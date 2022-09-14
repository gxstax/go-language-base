package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者（接收型 channel13）
func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer] send [%d] to channel13\n", i)
			i++
			continue
		}
		fmt.Printf("[producer] try send [%d], but channel13 is full\n", i)
	}
}

// 消费者（发送型 channel13）
func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Printf("[consumer]: try to recv from channel13, but the channel13 is empty\n")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: recy [%d] from channel13\n", i)
	}
}

// 发送条件判断
func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func main() {
	c := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}
