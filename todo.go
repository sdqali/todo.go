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
	foundItem, _ := repository.Find(item.id.String())
	fmt.Println(foundItem)
	_, err := repository.Find("some-id")
	fmt.Println(err)
}
