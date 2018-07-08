package main

type TodoStore interface {
	Add(todoItem TodoItem)
	Find(id string) (TodoItem, error)
	Remove(id string)
	All() []TodoItem
}
