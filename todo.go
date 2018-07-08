package main

import (
	"fmt"
)

func main() {
	store := InMemoryStore{}
	repository := TodoRepo{todoStore: &store}
	item := NewItem("Buy milk", false)
	repository.add(item)
	fmt.Println(repository)
}
