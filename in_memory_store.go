package todo

import "strings"

type InMemoryStore struct {
	items []TodoItem
}

func (store *InMemoryStore) Add(item TodoItem) {
	store.items = append(store.items, item)
}

func (store InMemoryStore) Get(id string) (TodoItem, error) {
	for _, item := range store.items {
		if item.Id.String() == id {
			return item, nil
		}
	}
	return TodoItem{}, TodoItemNotFound{possibleId: id}
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

func (store InMemoryStore) All() []TodoItem {
	return store.items
}

func (store *InMemoryStore) Save(itemToSave TodoItem) {
	for index, item := range store.items {
		if item.Id.String() == itemToSave.Id.String() {
			store.items[index] = itemToSave
			return
		}
	}
	store.Add(itemToSave)
}

func (store *InMemoryStore) Find(searchTerm string) []TodoItem {
	results := []TodoItem{}
	for _, item := range store.All() {
		if strings.Contains(strings.ToLower(item.Text), strings.ToLower(searchTerm)) {
			results = append(results, item)
		}
	}
	return results
}
