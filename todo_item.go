package main

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoItem struct {
	text string
	done bool
	id   uuid.UUID
}

func (item TodoItem) String() string {
	return fmt.Sprintf("%s, %s, %t", item.id, item.text, item.done)
}

func NewItem(text string, done bool) TodoItem {
	id, _ := uuid.NewRandom()
	return TodoItem{text: text, done: done, id: id}
}
