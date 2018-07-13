package todo

import "database/sql"

type DbStore struct {
	db *sql.DB
}

func NewDbStore(db *sql.DB) *DbStore {
	return &DbStore{db: db}
}

func (store *DbStore) Add(item TodoItem) {
	store.db.QueryRow("INSERT INTO todo_items(id, title, item_order, completed) VALUES($1, $2, $3, $4);", item.Id, item.Title, item.Order, item.Done)
}

func (store DbStore) Get(id string) (TodoItem, error) {
	return TodoItem{}, TodoItemNotFound{possibleId: id}
}

func (store *DbStore) Remove(id string) {
}

func (store DbStore) All() []TodoItem {
	var list []TodoItem
	return list
}

func (store *DbStore) Save(itemToSave TodoItem) {
}

func (store DbStore) Find(searchTerm string) []TodoItem {
	var list []TodoItem
	return list
}
