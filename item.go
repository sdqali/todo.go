package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type TodoItem struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Id        uuid.UUID `json:"id"`
	Order     int       `json:"order"`
}

func (item TodoItem) String() string {
	return fmt.Sprintf("%s, %s, %t", item.Id, item.Title, item.Completed)
}

func (item *TodoItem) MarkAsCompleted() {
	item.Completed = true
}

func (item *TodoItem) MarkAsTodo() {
	item.Completed = false
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
	return TodoItem{Title: text, Completed: false, Id: id}
}

func baseUrl() string {
	return os.Getenv("BASE_URL")
}
