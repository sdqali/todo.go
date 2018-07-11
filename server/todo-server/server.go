package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
	"github.com/sdqali/todo/server/todo-server/middleware"
	"github.com/sdqali/todo/server/todo-server/routes"
)

const filePath string = "/tmp/todo.json"

func main() {
	store := todo.NewJsonFileStore(filePath)
	repo := todo.NewTodoRepo(&store)
	router := mux.NewRouter()

	router.HandleFunc("/", routes.List(repo)).Methods("GET")
	router.HandleFunc("/", routes.Create(repo)).Methods("POST")
	router.HandleFunc("/", routes.Clear(repo)).Methods("DELETE")

	router.HandleFunc("/{id}", routes.Get(repo)).Methods("GET")
	router.HandleFunc("/{id}", routes.Patch(repo)).Methods("PATCH")
	router.HandleFunc("/{id}", routes.Delete(repo)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", middleware.WithMiddleWares(router)))
}
