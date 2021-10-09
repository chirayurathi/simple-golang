package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chirayurathi/task-app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dataBase string = "task-app-db"

var client *mongo.Client
var ctx context.Context

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func insertOne(doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection("users")
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "` + r.URL.Path[len("/users/"):] + `"}`))
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
		collection := client.Database(dataBase).Collection("user")
		result, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			// helper.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func handleRequests() {
	http.HandleFunc("/users/", addUser)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	var cancel context.CancelFunc
	var err error

	client, ctx, cancel, err = connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	ping(client, ctx)

	handleRequests()
}
