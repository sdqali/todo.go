package main

import "errors"

type InMemoryStore struct {
	items []TodoItem
}

func (store *InMemoryStore) add(todoItem TodoItem) {
	store.items = append(store.items, todoItem)
}

func (store InMemoryStore) find(id string) (TodoItem, error) {
	for _, item := range store.items {
		if item.id.String() == id {
			return item, nil
		}
	}
	return TodoItem{}, errors.New("Can't find item with given id")
}

func (store *InMemoryStore) remove(id string) {
	preserveIndex := 0
	for _, item := range store.items {
		if item.id.String() == id {
			store.items[preserveIndex] = item
			preserveIndex++
		}
	}
	store.items = store.items[:preserveIndex]
}

func (store InMemoryStore) all() []TodoItem {
	return store.items
}
