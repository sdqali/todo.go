package main

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TodoItem struct {
	text string
	done bool
	id   uuid.UUID
}

func (item TodoItem) String() string {
	return fmt.Sprintf("%s, %s, %t", item.id, item.text, item.done)
}

func NewItem(text string, done bool) TodoItem {
	id, _ := uuid.NewRandom()
	return TodoItem{text: text, done: done, id: id}
}

type TodoRepo struct {
	todoStore TodoStore
}

func (repo TodoRepo) String() string {
	var b bytes.Buffer
	for index, item := range repo.all() {
		b.WriteString(fmt.Sprintf("%d: %s", index, item.String()))
	}
	return b.String()
}

func (repo TodoRepo) all() []TodoItem {
	return repo.todoStore.all()
}

func (todoRepo TodoRepo) add(item TodoItem) {
	todoRepo.todoStore.add(item)
}

func (todoRepo TodoRepo) find(id string) (TodoItem, error) {
	return todoRepo.todoStore.find(id)
}

func (todoRepo TodoRepo) remove(id string) {
	todoRepo.todoStore.remove(id)
}

type TodoStore interface {
	add(todoItem TodoItem)
	find(id string) (TodoItem, error)
	remove(id string)
	all() []TodoItem
}

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

func main() {
	store := InMemoryStore{}
	repository := TodoRepo{todoStore: &store}
	item := NewItem("Buy milk", false)
	fmt.Println("Item: ", item)
	repository.add(item)
	fmt.Println(repository)
}
