package handlers

import (
	"net/http"
	"strconv"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
)

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	var page int
	var count int
	var err error
	id := r.URL.Path[len("/posts/users/"):]
	query := r.URL.Query()
	if len(query["page"]) > 0 && len(query["count"]) > 0 {
		page, err = strconv.Atoi(query["page"][0])
		if err != nil {
			helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		count, err = strconv.Atoi(query["count"][0])
		if err != nil {
			helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		page = -1
		count = -1
	}
	result, err := connection.GetAllPost(id, page, count)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteResponse(w, http.StatusOK, result)
}
