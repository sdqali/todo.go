package main

type TodoStore interface {
	Add(item TodoItem)
	Get(id string) (TodoItem, error)
	Remove(id string)
	All() []TodoItem
	Save(item TodoItem)
	Find(searchTerm string) []TodoItem
}
