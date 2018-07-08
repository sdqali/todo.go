package main

import (
	"bytes"
	"fmt"
)

type TodoRepo struct {
	todoStore TodoStore
}

func (repo TodoRepo) String() string {
	var b bytes.Buffer
	for index, item := range repo.All() {
		b.WriteString(fmt.Sprintf("%d: %s", index, item.String()))
	}
	return b.String()
}

func (repo TodoRepo) All() []TodoItem {
	return repo.todoStore.All()
}

func (repo TodoRepo) Add(item TodoItem) {
	repo.todoStore.Add(item)
}

func (repo TodoRepo) Find(id string) (TodoItem, error) {
	return repo.todoStore.Find(id)
}

func (repo TodoRepo) Remove(id string) {
	repo.todoStore.Remove(id)
}
