package todo

import "fmt"

type TodoItemNotFound struct {
	possibleId string
}

func (err TodoItemNotFound) Error() string {
	return fmt.Sprintf("Error: TodoItem with given id '%s' does not exist", err.possibleId)
}
