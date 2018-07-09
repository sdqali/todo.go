package todo

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type JsonFileStore struct {
	filePath string
}

func NewJsonFileStore(filePath string) JsonFileStore {
	return JsonFileStore{filePath: filePath}
}

func (store *JsonFileStore) Add(item TodoItem) {
	list := store.All()
	list = append(list, item)
	store.WriteRecord(list)
}

func (store JsonFileStore) Get(id string) (TodoItem, error) {
	for _, item := range store.All() {
		if item.Id.String() == id {
			return item, nil
		}
	}
	return TodoItem{}, TodoItemNotFound{possibleId: id}
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

func (store JsonFileStore) All() []TodoItem {
	bytes, _ := ioutil.ReadFile(store.filePath)
	var list []TodoItem
	json.Unmarshal(bytes, &list)
	return list
}

func (store *JsonFileStore) WriteRecord(items []TodoItem) {
	bytes, _ := json.Marshal(items)
	ioutil.WriteFile(store.filePath, bytes, 0644)
}

func (store *JsonFileStore) Save(itemToSave TodoItem) {
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

func (store JsonFileStore) Find(searchTerm string) []TodoItem {
	results := []TodoItem{}
	for _, item := range store.All() {
		if strings.Contains(strings.ToLower(item.Text), strings.ToLower(searchTerm)) {
			results = append(results, item)
		}
	}
	return results
}
