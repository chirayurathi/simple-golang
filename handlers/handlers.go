package handlers

import (
	"net/http"

	"github.com/chirayurathi/task-app/helpers"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		GetUser(w, r)
	case "POST":
		AddUser(w, r)
	default:
		helpers.WriteResponse(w, http.StatusMethodNotAllowed, `{"message": "Method Not Allowed"}`)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		GetPost(w, r)

	case "POST":
		AddPost(w, r)

	default:
		helpers.WriteResponse(w, http.StatusMethodNotAllowed, `{"message": "Method Not Allowed"}`)

	}
}

func UserPostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		GetUserPosts(w, r)

	default:
		helpers.WriteResponse(w, http.StatusMethodNotAllowed, `{"message": "Method Not Allowed"}`)

	}
}
