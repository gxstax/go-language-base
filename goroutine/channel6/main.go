package main

import (
	"fmt"
	"sync"
)

// 替代锁机制
type counter struct {
	ch chan int // 用户同步的 channel
	i  int      // 计数
}

func NewCounter() *counter {
	cter := &counter{
		ch: make(chan int),
	}

	go func() {
		// 死循环，ch 一旦可以发送，则立即发送
		for {
			cter.i++
			cter.ch <- cter.i
		}
	}()

	return cter
}

func (cter *counter) Increase() int {
	return <-cter.ch
}

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			n := counter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
