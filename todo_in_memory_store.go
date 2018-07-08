package main

type InMemoryStore struct {
	items []TodoItem
}

func (store *InMemoryStore) Add(todoItem TodoItem) {
	store.items = append(store.items, todoItem)
}

func (store InMemoryStore) Find(id string) (TodoItem, error) {
	for _, item := range store.items {
		if item.id.String() == id {
			return item, nil
		}
	}
	return TodoItem{}, TodoItemNotFound{possibleId: id}
}

func (store *InMemoryStore) Remove(id string) {
	preserveIndex := 0
	for _, item := range store.items {
		if item.id.String() == id {
			store.items[preserveIndex] = item
			preserveIndex++
		}
	}
	store.items = store.items[:preserveIndex]
}

func (store InMemoryStore) All() []TodoItem {
	return store.items
}
