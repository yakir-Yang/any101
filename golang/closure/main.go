package main

import (
	"fmt"
	"time"
)

func myfunc(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func myfunc2() []func() {
	var b []func()

	for i := 0; i < 3; i++ {
		b = append(b, func(j int) func() {
			return func() {
				fmt.Println("myfunc2: ", j)
			}
		}(i))
	}

	return b
}

func main() {
	i := 0

	defer func() {
		fmt.Println("main exit: ", i)
	}()

	test1 := myfunc(i)
	test2 := myfunc(i)

	test3 := func() int {
		i++
		return i
	}
	test4 := func() int {
		i++
		return i
	}

	fmt.Println("test1: ", test1())
	fmt.Println("test2: ", test2())
	fmt.Println("test3: ", test3())
	fmt.Println("test4: ", test4())

	atest := myfunc2()
	for _, f := range atest {
		f()
	}

	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}

	time.Sleep(time.Second)
}
