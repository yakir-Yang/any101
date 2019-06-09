// Package main provides ...
package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 4

func main() {
	ch := make(chan struct{}, 5)

	// 新建 cond
	var l sync.Mutex
	cond := sync.NewCond(&l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			// 条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Millisecond * 20)

	// 锁定一下
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}
}
