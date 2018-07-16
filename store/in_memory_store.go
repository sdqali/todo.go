package store

import (
	"strings"

	"github.com/sdqali/todo"
	"github.com/sdqali/todo/errors"
)

type InMemoryStore struct {
	items []todo.TodoItem
}

func (store *InMemoryStore) Add(item todo.TodoItem) {
	store.items = append(store.items, item)
}

func (store InMemoryStore) Get(id string) (todo.TodoItem, error) {
	for _, item := range store.items {
		if item.Id.String() == id {
			return item, nil
		}
	}
	return todo.TodoItem{}, errors.NotFound(id)
}

func (store *InMemoryStore) Remove(id string) {
	preserveIndex := 0
	for _, item := range store.items {
		if item.Id.String() != id {
			store.items[preserveIndex] = item
			preserveIndex++
		}
	}
	store.items = store.items[:preserveIndex]
}

func (store InMemoryStore) All() []todo.TodoItem {
	return store.items
}

func (store *InMemoryStore) Save(itemToSave todo.TodoItem) {
	for index, item := range store.items {
		if item.Id.String() == itemToSave.Id.String() {
			store.items[index] = itemToSave
			return
		}
	}
	store.Add(itemToSave)
}

func (store *InMemoryStore) Find(searchTerm string) []todo.TodoItem {
	results := []todo.TodoItem{}
	for _, item := range store.All() {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(searchTerm)) {
			results = append(results, item)
		}
	}
	return results
}
