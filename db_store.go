package todo

import (
	"database/sql"

	"github.com/google/uuid"
)

const SELECT_ALL_QUERY = "SELECT id, title, item_order, completed FROM todo_items;"
const SELECT_QUERY = "SELECT id, title, item_order, completed FROM todo_items WHERE id=$1;"
const INSERT_QUERY = "INSERT INTO todo_items(id, title, item_order, completed) VALUES($1, $2, $3, $4);"

type DbStore struct {
	db *sql.DB
}

func NewDbStore(db *sql.DB) *DbStore {
	return &DbStore{db: db}
}

func (store *DbStore) Add(item TodoItem) {
	store.db.QueryRow(INSERT_QUERY, item.Id, item.Title, item.Order, item.Done)
}

func (store DbStore) Get(id string) (TodoItem, error) {
	rows, _ := store.db.Query(SELECT_QUERY, id)
	list := listFromRows(rows)
	if len(list) == 0 {
		return TodoItem{}, TodoItemNotFound{possibleId: id}
	} else {
		return list[0], nil
	}
}

func listFromRows(rows *sql.Rows) []TodoItem {
	var list []TodoItem
	for rows.Next() {
		var id uuid.UUID
		var title string
		var order int
		var completed bool

		rows.Scan(&id, &title, &order, &completed)
		list = append(list, TodoItem{Id: id, Title: title, Order: order, Done: completed})
	}
	return list
}

func (store *DbStore) Remove(id string) {
}

func (store DbStore) All() []TodoItem {
	rows, _ := store.db.Query(SELECT_ALL_QUERY)
	return listFromRows(rows)
}

func (store *DbStore) Save(itemToSave TodoItem) {
}

func (store DbStore) Find(searchTerm string) []TodoItem {
	var list []TodoItem
	return list
}
