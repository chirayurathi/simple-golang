package handlers

import (
	"net/http"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	result, err := connection.GetUser(id)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	result.Password = string(helpers.Decrypt([]byte(result.Password)))
	helpers.WriteResponse(w, http.StatusOK, result)
}
