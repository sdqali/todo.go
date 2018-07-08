package main

type TodoStore interface {
	add(todoItem TodoItem)
	find(id string) (TodoItem, error)
	remove(id string)
	all() []TodoItem
}
