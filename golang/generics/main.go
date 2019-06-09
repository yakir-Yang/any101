package main

import (
	"fmt"
	"reflect"
)

// 类型断言实现泛型
func max1(first interface{}, rest ...interface{}) interface{} {
	m := first

	for _, v := range rest {
		switch v := v.(type) {
		case int:
			if v > m.(int) {
				m = v
			}
		case float32:
			if v > m.(float32) {
				m = v
			}
		default:
		}
	}

	return m
}

// 匿名接口实现泛型
type compareable interface {
	islarger(c interface{}) bool
}

func max2(first compareable, rest ...compareable) interface{} {
	m := first
	for _, v := range rest {
		if !m.islarger(v) {
			m = v
		}
	}
	return m
}

type mint int

func (m mint) islarger(s interface{}) bool {
	if m > s.(mint) {
		return true
	}
	return false
}

// 反射机制实现泛型
func max3(s interface{}) interface{} {
	if slice := reflect.ValueOf(s); slice.Kind() == reflect.Slice {
		m, ok := slice.Index(0).Interface().(float32)
		if !ok {
			fmt.Println("max3 only support float32 slices")
			return nil
		}

		for i := 0; i < slice.Len(); i++ {
			if m < slice.Index(i).Interface().(float32) {
				m = slice.Index(i).Interface().(float32)
			}
		}

		return m
	}

	fmt.Println("max3 only support float32 slices")
	return nil
}

func main() {
	fmt.Printf("类型断言实现泛型：max is %v\n", max1(1, 2, 5, 4, 3))
	fmt.Printf("匿名接口实现泛型：max is %v\n", max2(mint(1), mint(5), mint(4)))
	fmt.Printf("反射机制实现泛型：max is %v\n", max3([]float32{2.0, 1.0, 5.2}))
}
