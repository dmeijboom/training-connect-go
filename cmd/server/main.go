package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/dillendev/training-connect-go/gen/todo/v1/todov1connect"
)

func main() {
	service := &TodoService{}

	mux := http.NewServeMux()

	path, handler := todov1connect.NewTodoServiceHandler(service)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
