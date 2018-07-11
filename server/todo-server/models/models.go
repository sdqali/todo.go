package models

type TodoItemRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"completed"`
	Order int    `json:"order"`
}
