package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
}

type messageJSON struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

func main() {
	m1before := message{"Alice", "Hello", 12}
	m2before := messageJSON{"Alice", "Hello", 12}

	data1, _ := json.Marshal(m1before)
	data2, _ := json.Marshal(m2before)

	fmt.Println(string(data1))
	fmt.Println(string(data2))

	b := []byte(`{"Name":"Alice","Body":"Hello","Time":12}`)

	// should be True
	fmt.Println(bytes.Equal(data1, b))

	// should be False
	fmt.Println(bytes.Equal(data2, b))

	var m1after interface{}

	json.Unmarshal(data1, &m1after)

	fmt.Println(m1after)
}
