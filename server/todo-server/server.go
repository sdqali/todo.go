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
	router.HandleFunc("/{id}", routes.Get(repo)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8090", middleware.WithMiddleWares(router)))
}
