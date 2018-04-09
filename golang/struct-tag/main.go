package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string `昵称`
	sex  byte   `性别`
}

func main() {
	u := user{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}
