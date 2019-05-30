package main

import "fmt"

func getValue() interface{} {
	return 1
}

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

// This is type alias, T1 would also have m2 function
type T2 = T1

func (t T2) m2() {
	fmt.Println("T2.m1")
}

type MyStruct struct {
	T1
	k T2
}

func main() {
	i := getValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

	my := MyStruct{}
	my.m2()
}
