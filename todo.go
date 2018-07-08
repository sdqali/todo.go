package main

import (
	"fmt"
)

func main() {
	store := InMemoryStore{}
	repository := TodoRepo{todoStore: &store}
	item := NewItem("Buy milk")
	repository.Add(item)
	fmt.Println(repository)
	fmt.Println(repository.All())
	foundItem, _ := repository.Find(item.id.String())
	foundItem.MarkAsDone()
	repository.Save(foundItem)
	fmt.Println(repository.All())
	repository.Remove(item.id.String())
	fmt.Println(repository.All())
	_, err := repository.Find(item.id.String())
	fmt.Println(err)
}
