package main

import (
	"flag"
	"fmt"
	"strings"

	todo "github.com/sdqali/todo"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "/tmp/todo.json", "Path to a file to store Todo items in.")
	var action string
	flag.StringVar(&action, "action", "list", "An action to perform - one of list, add, get, mark-todo, mark-done, find, delete.")
	flag.Parse()
	store := todo.NewJsonFileStore(filePath)
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
		repo.MarkAsDone(id)
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
