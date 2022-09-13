package main

import (
	"sync"
	"time"
)

// 生产者消费者 channel 实现
func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
		if i == 5 {
			close(ch)
			return
		}
	}
	close(ch)
}

func consume(ch <-chan int) {
	for v := range ch {
		println(v)
	}
}
