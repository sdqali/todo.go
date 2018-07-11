package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type TodoItem struct {
	Title string    `json:"title"`
	Done  bool      `json:"completed"`
	Id    uuid.UUID `json:"id"`
}

func (item TodoItem) String() string {
	return fmt.Sprintf("%s, %s, %t", item.Id, item.Title, item.Done)
}

func (item *TodoItem) MarkAsDone() {
	item.Done = true
}

func (item *TodoItem) MarkAsTodo() {
	item.Done = false
}

func (item TodoItem) MarshalJSON() ([]byte, error) {
	type Alias TodoItem
	return json.Marshal(&struct {
		Url string `json:"url"`
		Alias
	}{
		Url:   fmt.Sprintf("%s/%s", baseUrl(), item.Id),
		Alias: (Alias)(item),
	})
}

func NewItem(text string) TodoItem {
	id, _ := uuid.NewRandom()
	return TodoItem{Title: text, Done: false, Id: id}
}

func baseUrl() string {
	return os.Getenv("BASE_URL")
}
