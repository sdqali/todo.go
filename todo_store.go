package main

type TodoStore interface {
	Add(item TodoItem)
	Find(id string) (TodoItem, error)
	Remove(id string)
	All() []TodoItem
	Save(item TodoItem)
}
