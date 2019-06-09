package main

import (
	"fmt"
	"reflect"
)

func index(s interface{}, v interface{}) int {
	if rs := reflect.ValueOf(s); rs.Kind() == reflect.Slice {
		for i := 0; i < rs.Len(); i++ {
			if reflect.DeepEqual(v, rs.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Value 3 at Index: ", index(s1, 3))
}
