package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
	"github.com/sdqali/todo/db"
	"github.com/sdqali/todo/server/todo-server/middleware"
	"github.com/sdqali/todo/server/todo-server/routes"
)

const filePath string = "/tmp/todo.json"

func main() {
	repo := getRepo()
	router := mux.NewRouter()

	router.HandleFunc("/", routes.List(repo)).Methods("GET")
	router.HandleFunc("/", routes.Create(repo)).Methods("POST")
	router.HandleFunc("/", routes.Clear(repo)).Methods("DELETE")

	router.HandleFunc("/{id}", routes.Get(repo)).Methods("GET")
	router.HandleFunc("/{id}", routes.Patch(repo)).Methods("PATCH")
	router.HandleFunc("/{id}", routes.Delete(repo)).Methods("DELETE")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, middleware.WithMiddleWares(router)))
}

func getRepo() todo.TodoRepo {
	return todo.NewTodoRepo(getStore())
}

func getStore() todo.TodoStore {
	var storeType string
	flag.StringVar(&storeType, "store", "in-memory", "One of json-file or in-memory")
	flag.Parse()

	var store todo.TodoStore

	switch storeType {
	case "json-file":
		store = todo.NewJsonFileStore(filePath)
	case "in-memory":
		store = &todo.InMemoryStore{}
	case "db":
		db := db.GetDb()
		store = todo.NewDbStore(db)
	default:
		store = &todo.InMemoryStore{}
	}

	return store
}
