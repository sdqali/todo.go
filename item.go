package todo

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoItem struct {
	Text string    `json:"title"`
	Done bool      `json:"completed"`
	Id   uuid.UUID `json:"id"`
	Url  string    `json:"url"`
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

func NewItem(text string, baseUrl string) TodoItem {
	id, _ := uuid.NewRandom()
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	return TodoItem{Text: text, Done: false, Id: id, Url: url}
}
