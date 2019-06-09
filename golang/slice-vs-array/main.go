package main

import (
	"fmt"
	"reflect"
)

// passing array as argument
func arraytest(x [3]int) {
	fmt.Printf("-> arraytest: x = %v, &x = [%p, %v]\n", x, &x, reflect.TypeOf(x))
	fmt.Printf("-> arraytest: &x[0] = %p\n", &x[0])
	x[1] += 100
}

// passing slice as argument
func slicetest(x []int) {
	fmt.Printf("-> slicetest: x = %v, &x = [%p, %v]\n", x, &x, reflect.TypeOf(x))
	fmt.Printf("-> slicetest: &x[0] = %p\n", &x[0])

	x[1] += 100
}

func main() {
	s1 := make([]int, 3)
	s2 := [3]int{1, 2, 3}

	fmt.Printf("s1 = %v, &s1 = [%p, %v]\n", s1, &s1, reflect.TypeOf(s1))
	fmt.Printf("&s1[0] = %p\n", &s1[0])
	slicetest(s1)
	fmt.Printf("slicetest(s1) = %v [%p]\n", s1, &s1)

	fmt.Printf("\n")

	fmt.Printf("s2 = %v, &s2 = [%p, %v]\n", s2, &s2, reflect.TypeOf(s1))
	fmt.Printf("&s2[0] = %p\n", &s2[0])
	arraytest(s2)
	fmt.Printf("arraytest(s2) = %v [%p]\n", s2, &s2)

	fmt.Printf("\n")

	// reslice would affect origin slice
	s3 := [...]int{1, 2, 3, 4, 5, 6}
	res := s3[2:4]
	res = append(res, 123)
	fmt.Println(res, s3)
}
