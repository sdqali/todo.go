package main

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoItem struct {
	Text string    `json:"text"`
	Done bool      `json:"done"`
	Id   uuid.UUID `json:"id"`
}

func (item TodoItem) String() string {
	return fmt.Sprintf("%s, %s, %t", item.Id, item.Text, item.Done)
}

func (item *TodoItem) MarkAsDone() {
	item.Done = true
}

func (item *TodoItem) MarkAsTodo() {
	item.Done = false
}

func NewItem(text string) TodoItem {
	id, _ := uuid.NewRandom()
	return TodoItem{Text: text, Done: false, Id: id}
}
