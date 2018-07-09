package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
)

const filePath string = "/tmp/todo.json"

func main() {
	store := todo.NewJsonFileStore(filePath)
	repo := todo.NewTodoRepo(&store)
	router := mux.NewRouter()

	router.HandleFunc("/", List(repo)).Methods("GET")

	router.HandleFunc("/{id}", Get(repo)).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func List(repo todo.TodoRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repo.All())
	}
}

func Get(repo todo.TodoRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		item, err := repo.Get(vars["id"])
		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
