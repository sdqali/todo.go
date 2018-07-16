package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
	"github.com/sdqali/todo/server/todo-server/middleware"
	"github.com/sdqali/todo/server/todo-server/routes"
	st "github.com/sdqali/todo/store"
	cassandra "github.com/sdqali/todo/store/cassandra"
	js "github.com/sdqali/todo/store/json"
	pg "github.com/sdqali/todo/store/postgres"
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
		store = js.NewJsonFileStore(filePath)
	case "in-memory":
		store = &st.InMemoryStore{}
	case "pg":
		db := pg.GetDb()
		store = pg.NewDbStore(db)
	case "cassandra":
		cluster := cassandra.GetCluster()
		store = cassandra.NewCassandraStore(cluster)
	default:
		store = &st.InMemoryStore{}
	}

	return store
}
