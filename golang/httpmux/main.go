package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var todos Todos

func init() {
	todos = Todos{
		Todo{
			Name:      "Write presentation",
			Completed: false,
			Due:       time.Now(),
		},
		Todo{
			Name:      "Host meetup",
			Completed: true,
			Due:       time.Now(),
		},
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", logger(index))
	router.Handle("/todos", logger(todoIndex))
	router.Handle("/todos/{todoId}", logger(todoShow))

	log.Println("curl http://localhost:8001/todos/Write%20presentation | jq")
	log.Println("curl http://localhost:8001/todos | jq")
	log.Println("curl http://localhost:8001")

	log.Fatal(http.ListenAndServe(":8001", router))
}

func logger(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		f(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]

	var todo Todo

	for _, todo = range todos {
		if todo.Name == todoID {
			break
		}
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}
