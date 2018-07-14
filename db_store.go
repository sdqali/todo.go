package todo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const SELECT_ALL_QUERY = "SELECT id, title, item_order, completed FROM todo_items;"
const SELECT_QUERY = "SELECT id, title, item_order, completed FROM todo_items WHERE id=$1 LIMIT 1;"
const INSERT_QUERY = "INSERT INTO todo_items(id, title, item_order, completed) VALUES($1, $2, $3, $4);"
const DELETE_QUERY = "DELETE FROM todo_items WHERE id=$1;"
const UPDATE_QUERY = "UPDATE todo_items SET title=$1, item_order=$2, completed=$3 WHERE id=$4;"
const SEARCH_QUERY = "SELECT id, title, item_order, completed FROM todo_items WHERE title ILIKE '%%%s%%';"

type DbStore struct {
	db *sql.DB
}

func NewDbStore(db *sql.DB) *DbStore {
	return &DbStore{db: db}
}

func (store *DbStore) Add(item TodoItem) {
	rows, _ := store.db.Query(INSERT_QUERY, item.Id, item.Title, item.Order, item.Done)
	rows.Close()
}

func (store DbStore) Get(id string) (TodoItem, error) {
	rows, err := store.db.Query(SELECT_QUERY, id)
	defer rows.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return TodoItem{}, TodoItemNotFound{possibleId: id}
	}
	list := itemsFromRows(rows)
	if len(list) == 0 {
		return TodoItem{}, TodoItemNotFound{possibleId: id}
	} else {
		return list[0], nil
	}
}

func (store *DbStore) Remove(id string) {
	rows, err := store.db.Query(DELETE_QUERY, id)
	defer rows.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (store DbStore) All() []TodoItem {
	rows, err := store.db.Query(SELECT_ALL_QUERY)
	defer rows.Close()
	if err == nil {
		return itemsFromRows(rows)
	} else {
		fmt.Println("ERROR: ", err)
		return make([]TodoItem, 0)
	}
}

func (store *DbStore) Save(itemToSave TodoItem) {
	rows, _ := store.db.Query(UPDATE_QUERY, itemToSave.Title, itemToSave.Order, itemToSave.Done, itemToSave.Id)
	rows.Close()
}

func (store DbStore) Find(searchTerm string) []TodoItem {
	query := fmt.Sprintf(SEARCH_QUERY, searchTerm)
	rows, err := store.db.Query(query)
	defer rows.Close()
	if err == nil {
		return itemsFromRows(rows)
	} else {
		fmt.Println("ERROR: ", err)
		return make([]TodoItem, 0)
	}
}

func itemsFromRows(rows *sql.Rows) []TodoItem {
	list := make([]TodoItem, 0)
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
