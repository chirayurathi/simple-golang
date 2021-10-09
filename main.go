package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chirayurathi/task-app/connection"
	"github.com/chirayurathi/task-app/helpers"
	"github.com/chirayurathi/task-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func addUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		id := r.URL.Path[len("/users/"):]
		result, err := connection.GetUser(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		result.Password = string(helpers.Decrypt([]byte(result.Password)))
		WriteResponse(w, http.StatusOK, result)

	case "POST":
		user := models.User{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = json.Unmarshal(body, &user)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		user.Password = string(helpers.Encrypt([]byte(user.Password)))
		result, err := connection.InsertUser(user)

		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		id := r.URL.Path[len("/posts/"):]
		result, err := connection.GetPost(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, result)

	case "POST":
		post := models.Post{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &post)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		post.Posted_timestamp = primitive.Timestamp{T: uint32(time.Now().Unix())}

		result, err := connection.InsertPost(post)

		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func userPostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		var page int
		var count int
		var err error
		id := r.URL.Path[len("/posts/users/"):]
		query := r.URL.Query()
		if len(query["page"]) > 0 && len(query["count"]) > 0 {
			page, err = strconv.Atoi(query["page"][0])
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			count, err = strconv.Atoi(query["count"][0])
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			page = -1
			count = -1
		}
		result, err := connection.GetAllPost(id, page, count)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteResponse(w, http.StatusOK, result)
	}
}

func handleRequests() {
	http.HandleFunc("/users/", addUser)
	http.HandleFunc("/posts/", addPost)
	http.HandleFunc("/posts/users/", userPostsHandler)
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
