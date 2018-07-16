package todo

import (
	"bytes"
	"fmt"

	"github.com/sdqali/todo/domain"
	"github.com/sdqali/todo/store"
)

type TodoRepo struct {
	store store.TodoStore
}

func NewTodoRepo(store store.TodoStore) TodoRepo {
	return TodoRepo{store: store}
}

func (repo TodoRepo) String() string {
	var b bytes.Buffer
	for index, item := range repo.All() {
		b.WriteString(fmt.Sprintf("%d: %s\n", index, item.String()))
	}
	return b.String()
}

func (repo TodoRepo) All() []domain.TodoItem {
	return repo.store.All()
}

func (repo TodoRepo) Add(item domain.TodoItem) {
	repo.store.Add(item)
}

func (repo TodoRepo) Get(id string) (domain.TodoItem, error) {
	return repo.store.Get(id)
}

func (repo TodoRepo) Remove(id string) {
	repo.store.Remove(id)
}

func (repo TodoRepo) Clear() {
	for _, item := range repo.store.All() {
		repo.Remove(item.Id.String())
	}
}

func (repo TodoRepo) Save(item domain.TodoItem) {
	repo.store.Save(item)
}

func (repo TodoRepo) MarkAsCompleted(id string) {
	item, err := repo.Get(id)
	if err == nil {
		item.MarkAsCompleted()
		fmt.Println(item)
		repo.Save(item)
	}
}

func (repo TodoRepo) MarkAsTodo(id string) {
	item, err := repo.Get(id)
	if err == nil {
		item.MarkAsTodo()
		fmt.Println(item)
		repo.Save(item)
	}
}

func (repo TodoRepo) Find(searchTerm string) []domain.TodoItem {
	return repo.store.Find(searchTerm)
}
