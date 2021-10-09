package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
	"github.com/chirayurathi/task-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	post.Posted_timestamp = primitive.Timestamp{T: uint32(time.Now().Unix())}

	result, err := connection.InsertPost(post)

	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(result)
}
