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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repo.All())
	})

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		item, err := repo.Get(vars["id"])
		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
