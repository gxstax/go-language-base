package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: is works done!\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			// 所有工作现成阻塞在组信号 groupSignal 上
			<-groupSignal
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

// 无缓冲 channel13 用来实现 1 对 n 的信号通知机制
func main() {
	// 创建组信号 channel13
	groupSignal := make(chan signal)
	// 启动工作现成
	c := spawnGroup(worker, 5, groupSignal)

	time.Sleep(3 * time.Second)
	fmt.Println("the group of workers start to worker...")
	// 关闭 groupSignal 使得阻塞线程得以执行
	close(groupSignal)

	// 等待所有 worker 线程执行完成
	<-c
	fmt.Println("the group of workers work done!")

}
