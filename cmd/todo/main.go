package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sdqali/todo"
	pg "github.com/sdqali/todo/db/postgres"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "/tmp/todo.json", "Path to a file to store Todo items in.")
	var action string
	flag.StringVar(&action, "action", "list", "An action to perform - one of list, add, get, mark-todo, mark-done, find, delete.")
	var storeType string
	flag.StringVar(&storeType, "store", "in-memory", "One of json-file or in-memory")
	flag.Parse()

	var store todo.TodoStore

	switch storeType {
	case "json-file":
		store = todo.NewJsonFileStore(filePath)
	case "in-memory":
		store = &todo.InMemoryStore{}
	case "pg":
		db := pg.GetDb()
		store = todo.NewDbStore(db)
	default:
		store = &todo.InMemoryStore{}
	}

	repo := todo.NewTodoRepo(store)
	switch action {
	case "list":
		fmt.Println(repo)
	case "add":
		text := strings.Join(flag.Args(), " ")
		repo.Add(todo.NewItem(text))
	case "get":
		id := flag.Args()[0]
		item, err := repo.Get(id)
		if err == nil {
			fmt.Println(item)
		} else {
			fmt.Println(err)
		}
	case "mark-done":
		id := flag.Args()[0]
		repo.MarkAsCompleted(id)
	case "mark-todo":
		id := flag.Args()[0]
		repo.MarkAsTodo(id)
	case "delete":
		id := flag.Args()[0]
		repo.Remove(id)
	case "find":
		id := flag.Args()[0]
		for _, item := range repo.Find(id) {
			fmt.Println(item)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
