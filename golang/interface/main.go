package main

import (
	"fmt"
)

type people interface {
	Show()
}

type student struct{}

func (stu *student) Show() {

}

//func live() interface{} {
func live() people {
	var stu *student
	//var stu interface{}
	//var stu *int
	if stu == nil {
		fmt.Println("stu: ", stu)
	}
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
