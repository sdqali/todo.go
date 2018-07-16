package json

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/sdqali/todo/domain"
	"github.com/sdqali/todo/errors"
)

type JsonFileStore struct {
	filePath string
}

func NewJsonFileStore(filePath string) *JsonFileStore {
	return &JsonFileStore{filePath: filePath}
}

func (store *JsonFileStore) Add(item domain.TodoItem) {
	list := store.All()
	list = append(list, item)
	store.WriteRecord(list)
}

func (store JsonFileStore) Get(id string) (domain.TodoItem, error) {
	for _, item := range store.All() {
		if item.Id.String() == id {
			return item, nil
		}
	}
	return domain.TodoItem{}, errors.NotFound(id)
}

func (store *JsonFileStore) Remove(id string) {
	preserveIndex := 0
	allItems := store.All()
	for _, item := range store.All() {
		if item.Id.String() != id {
			allItems[preserveIndex] = item
			preserveIndex++
		}
	}
	store.WriteRecord(allItems[:preserveIndex])
}

func (store JsonFileStore) All() []domain.TodoItem {
	bytes, _ := ioutil.ReadFile(store.filePath)
	var list []domain.TodoItem
	json.Unmarshal(bytes, &list)
	return list
}

func (store *JsonFileStore) WriteRecord(items []domain.TodoItem) {
	bytes, _ := json.Marshal(items)
	ioutil.WriteFile(store.filePath, bytes, 0644)
}

func (store *JsonFileStore) Save(itemToSave domain.TodoItem) {
	allItems := store.All()
	for index, item := range allItems {
		if item.Id.String() == itemToSave.Id.String() {
			allItems[index] = itemToSave
			store.WriteRecord(allItems)
			return
		}
	}
	store.Add(itemToSave)
}

func (store JsonFileStore) Find(searchTerm string) []domain.TodoItem {
	results := []domain.TodoItem{}
	for _, item := range store.All() {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(searchTerm)) {
			results = append(results, item)
		}
	}
	return results
}
