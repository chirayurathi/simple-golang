package handlers

import (
	"net/http"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]
	result, err := connection.GetPost(id)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteResponse(w, http.StatusOK, result)
}
