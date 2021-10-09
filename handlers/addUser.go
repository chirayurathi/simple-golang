package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
	"github.com/chirayurathi/task-app/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(user.Name) == 0 {
		helpers.WriteResponse(w, http.StatusBadRequest, `{message:"This Field Is Required"}`)
		return
	}
	if len(user.Email) == 0 || !helpers.ValidEmail(user.Email) {
		helpers.WriteResponse(w, http.StatusBadRequest, `{message:"Enter Valid Email"}`)
		return
	}
	if len(user.Password) == 0 {
		helpers.WriteResponse(w, http.StatusBadRequest, `{message:"This Field Is Required"}`)
		return
	}
	user.Password = string(helpers.Encrypt([]byte(user.Password)))
	result, err := connection.InsertUser(user)

	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(result)
}
