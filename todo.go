package main

import (
	"fmt"
)

func main() {
	store := InMemoryStore{}
	repository := TodoRepo{todoStore: &store}
	item := NewItem("Buy milk", false)
	repository.Add(item)
	fmt.Println(repository)
	fmt.Println(repository.All())
	foundItem, _ := repository.Find(item.id.String())
	fmt.Println(foundItem)
	repository.Remove(item.id.String())
	fmt.Println(repository.All())
	_, err := repository.Find(item.id.String())
	fmt.Println(err)
}
