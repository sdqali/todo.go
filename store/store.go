package store

import "github.com/sdqali/todo/domain"

type TodoStore interface {
	Add(item domain.TodoItem)
	Get(id string) (domain.TodoItem, error)
	Remove(id string)
	All() []domain.TodoItem
	Save(item domain.TodoItem)
	Find(searchTerm string) []domain.TodoItem
}
