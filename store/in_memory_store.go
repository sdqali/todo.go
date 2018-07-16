package store

import (
	"strings"

	"github.com/sdqali/todo/domain"
	"github.com/sdqali/todo/errors"
)

type InMemoryStore struct {
	items []domain.TodoItem
}

func (store *InMemoryStore) Add(item domain.TodoItem) {
	store.items = append(store.items, item)
}

func (store InMemoryStore) Get(id string) (domain.TodoItem, error) {
	for _, item := range store.items {
		if item.Id.String() == id {
			return item, nil
		}
	}
	return domain.TodoItem{}, errors.NotFound(id)
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

func (store InMemoryStore) All() []domain.TodoItem {
	return store.items
}

func (store *InMemoryStore) Save(itemToSave domain.TodoItem) {
	for index, item := range store.items {
		if item.Id.String() == itemToSave.Id.String() {
			store.items[index] = itemToSave
			return
		}
	}
	store.Add(itemToSave)
}

func (store *InMemoryStore) Find(searchTerm string) []domain.TodoItem {
	results := []domain.TodoItem{}
	for _, item := range store.All() {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(searchTerm)) {
			results = append(results, item)
		}
	}
	return results
}
