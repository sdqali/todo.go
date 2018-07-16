package errors

import "fmt"

type TodoItemNotFound struct {
	PossibleId string
}

func (err TodoItemNotFound) Error() string {
	return fmt.Sprintf("Error: TodoItem with given id '%s' does not exist", err.PossibleId)
}

func NotFound(id string) TodoItemNotFound {
	return TodoItemNotFound{PossibleId: id}
}
