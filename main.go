package main

import (
	"log"
	"net/http"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/handlers"
)

func handleRequests() {
	http.HandleFunc("/users/", handlers.UserHandler)
	http.HandleFunc("/posts/", handlers.PostHandler)
	http.HandleFunc("/posts/users/", handlers.UserPostsHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	connection.Connect("mongodb://127.0.0.1:27017/")
	if connection.Err != nil {
		panic(connection.Err)
	}

	// defer connection.Close(connection.Client, connection.Ctx, connection.Cancel)

	connection.Ping(connection.Client, connection.Ctx)

	handleRequests()
}
