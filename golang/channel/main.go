package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！
	//ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, _ := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	time.Sleep(100 * time.Millisecond)

	select {
	case i := <-c:
		fmt.Println("select: ", i)
	default:
		fmt.Println("default")
	}

	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println("range: ", i)
	}

	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4", "5"},
	}
	s := <-th.Iter()
	fmt.Printf("%s %v\n", "ch", s)
}
