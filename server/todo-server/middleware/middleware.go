package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func WithMiddleWares(router *mux.Router) http.Handler {
	return JsonMiddleWare(WithCors(router))
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
