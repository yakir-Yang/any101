package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

/*
 * Handle is my private handle
 */
type Handle struct{}

/*
 * handle the resouce action
 */
func (h *Handle) Test1Action(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test1 Action")
	w.Write(([]byte("Test1")))
}

func (h *Handle) Test2Action(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test2 Action")
	w.Write(([]byte("Test2")))
}

func myhandle(w http.ResponseWriter, r *http.Request) {
	pathinfo := strings.Trim(r.URL.Path, "/")
	paths := strings.Split(pathinfo, "/")

	fmt.Println("===> ", r.URL)
	fmt.Println("paths: ", strings.Join(paths, "|"))

	action := ""
	if len(paths) > 1 {
		action = strings.Title(paths[1] + "Action")
	}

	fmt.Println("actionfunc: ", action)

	handle := &Handle{}

	controller := reflect.ValueOf(handle)

	method := controller.MethodByName(action)

	nr := reflect.ValueOf(r)
	nw := reflect.ValueOf(w)

	method.Call([]reflect.Value{nw, nr})
}

func main() {
	// directly respond to hello website
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		message = strings.TrimPrefix(message, "/")
		message = "Welcome to my website: " + message

		w.Write([]byte(message))
	})

	// find the method by path name
	http.Handle("/handle/", http.HandlerFunc(myhandle))

	fmt.Println("http server listen at http://localhost:8001")

	fmt.Println("Avaliable path:")
	fmt.Println("  - /hello/")
	fmt.Println("  - /handle/test1")
	fmt.Println("  - /handle/test2")

	if err := http.ListenAndServe(":8001", nil); err != nil {
		fmt.Printf("failed to listen at :8081 port")
		panic(err)
	}
}
