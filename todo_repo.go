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
