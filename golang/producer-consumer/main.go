package main

import (
	"fmt"
	"sync"
)

func producer(wg *sync.WaitGroup, ch chan<- int, id int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("producer[%d]: %d\n", id, i)
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch <-chan int, id int) {
	for i := range ch {
		fmt.Printf("consumer[%d]: %d\n", id, i)
	}
	wg.Done()
}

func main() {
	var wgpro sync.WaitGroup
	var wgcon sync.WaitGroup

	ch := make(chan int, 10)

	for i := 0; i < 2; i++ {
		wgpro.Add(1)
		go producer(&wgpro, ch, i)

		wgcon.Add(1)
		go consumer(&wgcon, ch, i)
	}

	wgpro.Wait()

	close(ch)

	wgcon.Wait()
}
