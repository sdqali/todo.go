package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
	"github.com/sdqali/todo/server/todo-server/models"
)

func Create(repo todo.TodoRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var itemRequest models.TodoItemRequest
		bytes, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(bytes, &itemRequest)
		item := todo.NewItem(itemRequest.Title)
		repo.Add(item)
		json.NewEncoder(writer).Encode(item)
	}
}

func List(repo todo.TodoRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(repo.All())
	}
}

func Get(repo todo.TodoRepo) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		item, err := repo.Get(vars["id"])
		if err == nil {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(item)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}
