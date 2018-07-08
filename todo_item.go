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

func (item *TodoItem) MarkAsDone() {
	item.done = true
}

func (item *TodoItem) MarkAsTodo() {
	item.done = false
}

func NewItem(text string) TodoItem {
	id, _ := uuid.NewRandom()
	return TodoItem{text: text, done: false, id: id}
}
