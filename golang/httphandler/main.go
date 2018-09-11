package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
)

type (
	fruitMap map[string]interface{}
)

func fruitsHandler() http.Handler {
	fruits := fruitMap{}

	mux := http.NewServeMux()

	mux.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		handleFruitList(fruits, w, r)
	})

	mux.HandleFunc("/fruits/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		handleFruit(fruits, w, r)
	})

	return mux
}

func handleFruitList(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ret := []string{}
		for k := range fruits {
			ret = append(ret, k)
		}

		b, err := json.Marshal(ret)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleFruit(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	_, name := path.Split(r.URL.Path)

	switch r.Method {
	case "GET":
		if data, ok := fruits[name]; ok {
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	case "PUT":
		var data map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			panic(err)
		}

		fruits[name] = data
		w.WriteHeader(http.StatusNoContent)

	case "DELETE":
		if _, ok := fruits[name]; ok {
			delete(fruits, name)
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	handler := fruitsHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	fmt.Println("Server listen at ", server.URL)

	select {}
}
