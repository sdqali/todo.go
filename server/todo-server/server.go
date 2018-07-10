package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
	"github.com/sdqali/todo/server/todo-server/middleware"
)

const filePath string = "/tmp/todo.json"

func main() {
	store := todo.NewJsonFileStore(filePath)
	repo := todo.NewTodoRepo(&store)
	router := mux.NewRouter()

	router.HandleFunc("/", List(repo)).Methods("GET")
	router.HandleFunc("/{id}", Get(repo)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8090", middleware.WithMiddleWares(router)))
}

func List(repo todo.TodoRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(repo.All())
	}
}

func Get(repo todo.TodoRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		item, err := repo.Get(vars["id"])
		if err == nil {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(item)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}
