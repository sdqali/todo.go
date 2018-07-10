package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sdqali/todo"
)

const filePath string = "/tmp/todo.json"

func main() {
	store := todo.NewJsonFileStore(filePath)
	repo := todo.NewTodoRepo(&store)
	router := mux.NewRouter()

	router.HandleFunc("/", List(repo)).Methods("GET")

	router.HandleFunc("/{id}", Get(repo)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", JsonMiddleWare(WithCors(router))))
}

func JsonMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(writer, request)
	})
}

func WithCors(router *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("CORS_ALLOWED_ORIGINS")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
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
