package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle := new(cgi.Handler)

		handle.Path = "/usr/local/bin/go"

		pwd, _ := os.Getwd()
		handle.Dir = pwd + "/cgi-scripts"

		script := handle.Dir + r.URL.Path

		args := []string{"run", script}
		handle.Args = append(handle.Args, args...)

		fmt.Println("handlePath: ", handle.Path)
		fmt.Println("handleArgs: ", handle.Args)

		handle.ServeHTTP(w, r)
	})

	fmt.Println("CGI server listen at http://localhost:8001")

	fmt.Println("Avaliable cgi files:")

	files, _ := filepath.Glob("cgi-scripts/*")
	fmt.Println(files)

	http.ListenAndServe(":8001", nil)
}
