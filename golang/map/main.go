package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
 * MapTest Nothing comment
 */
type MapTest struct {
	imap map[string]int
	sync.RWMutex
}

func NewMapTest() *MapTest {
	var mt MapTest
	mt.imap = make(map[string]int)
	return &mt
}

func (mt *MapTest) Add(key string, val int) {
	mt.Lock()
	defer mt.Unlock()
	mt.imap[key] = val
}

func (mt *MapTest) Get(key string) int {
	mt.RLock()
	defer mt.RUnlock()
	if val, ok := mt.imap[key]; ok {
		return val
	}
	return -1
}

func main() {
	test := make(map[int]int)

	test[0] = 0

	for key, val := range test {
		fmt.Printf("key: %d, val: %d\n", key, val)
	}

	mt := NewMapTest()

	for i := 0; i < 1000; i++ {
		go func(id int) {
			mt.Add(strconv.Itoa(id), id)
		}(i)
	}

	for i := 0; i < 1000; i++ {
		go func(id int) {
			//fmt.Printf("Get: %d\n", mt.Get(strconv.Itoa(id)))
			mt.Get(strconv.Itoa(id))
		}(i)
	}

	fmt.Printf("Test Over\n")

	for {
	}
}
