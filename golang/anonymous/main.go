package main

import "fmt"

type attr struct {
	perm int
	int  // size
}

type file struct {
	name string
	attr
}

func main() {
	d := file{
		name: "test.attr",
		attr: attr{
			perm: 0755,
			int:  1024,
		},
	}

	fmt.Println(d)
}
