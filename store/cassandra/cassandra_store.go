package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/sdqali/todo/domain"
	"github.com/sdqali/todo/errors"
)

const SELECT_ALL_QUERY = "SELECT id, title, item_order, completed FROM todo_items;"
const INSERT_QUERY = "INSERT INTO todo_items(id, title, item_order, completed) VALUES(?, ?, ?, ?);"
const SELECT_QUERY = "SELECT id, title, item_order, completed FROM todo_items WHERE id=? LIMIT 1;"
const DELETE_QUERY = "DELETE FROM todo_items WHERE id=?;"
const UPDATE_QUERY = "UPDATE todo_items SET title=?, item_order=?, completed=? WHERE id=?;"
const SEARCH_QUERY = "SELECT id, title, item_order, completed FROM todo_items WHERE title LIKE '%%%s%%' ALLOW FILTERING;"

type CassandraStore struct {
	cluster *gocql.ClusterConfig
}

func NewCassandraStore(cluster *gocql.ClusterConfig) *CassandraStore {
	return &CassandraStore{cluster: cluster}
}

func (store *CassandraStore) Add(item domain.TodoItem) {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	err = session.Query(INSERT_QUERY, item.Id.String(), item.Title, item.Order, item.Completed).Exec()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (store *CassandraStore) Get(id string) (domain.TodoItem, error) {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return domain.TodoItem{}, errors.NotFound(id)
	}
	iter := session.Query(SELECT_QUERY, id).Iter()
	items := itemsFromIter(iter)
	if len(items) == 0 {
		return domain.TodoItem{}, errors.NotFound(id)
	} else {
		return items[0], nil
	}
}

func (store *CassandraStore) Remove(id string) {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	err = session.Query(DELETE_QUERY, id).Exec()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (store *CassandraStore) All() []domain.TodoItem {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return make([]domain.TodoItem, 0)
	}
	iter := session.Query(SELECT_ALL_QUERY).Iter()
	return itemsFromIter(iter)
}

func (store *CassandraStore) Save(item domain.TodoItem) {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	err = session.Query(UPDATE_QUERY, item.Title, item.Order, item.Completed, item.Id.String()).Exec()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (store *CassandraStore) Find(searchTerm string) []domain.TodoItem {
	session, err := store.cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return make([]domain.TodoItem, 0)
	}
	query := fmt.Sprintf(SEARCH_QUERY, searchTerm)
	iter := session.Query(query).Iter()
	return itemsFromIter(iter)
}

func itemsFromIter(iter *gocql.Iter) []domain.TodoItem {
	list := make([]domain.TodoItem, 0)

	var id_str string
	var title string
	var order int
	var completed bool

	for iter.Scan(&id_str, &title, &order, &completed) {
		id, _ := uuid.Parse(id_str)
		list = append(list, domain.TodoItem{Id: id, Title: title, Order: order, Completed: completed})
	}
	return list
}
