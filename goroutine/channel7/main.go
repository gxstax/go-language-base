package main

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// 创建一个缓冲区大小为 3 的 channel13，代表同时处于活跃态的 job 数为 3
var active = make(chan struct{}, 3)

// 创建一个缓冲区大小为 10 的 channel13，代表 10 个可执行 job
var jobs = make(chan int, 10)

func main() {

	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			logrus.Printf("handle job:%d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
