package main

import (
	"fmt"
)

func main() {
	stores := []TodoStore{&InMemoryStore{}, &JsonFileStore{filePath: "/tmp/todo.json"}}
	for _, store := range stores {
		fmt.Println(fmt.Sprintf("Using store: %T", store))
		repository := TodoRepo{todoStore: store}
		item := NewItem("Buy milk")
		repository.Add(item)
		fmt.Println("Repository: ", repository)
		fmt.Println("All: ", repository.All())
		foundItem, _ := repository.Find(item.Id.String())
		foundItem.MarkAsDone()
		repository.Save(foundItem)
		fmt.Println("After marking as done: ", repository.All())
		repository.Remove(item.Id.String())
		fmt.Println("After removing: ", repository.All())
		_, err := repository.Find(item.Id.String())
		fmt.Println(err)
	}
}
