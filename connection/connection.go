package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DataBase string = "task-app-db"

var Client *mongo.Client
var Ctx context.Context
var Cancel context.CancelFunc
var Err error

func Close(Client *mongo.Client, Ctx context.Context, Cancel context.CancelFunc) {

	defer Cancel()

	defer func() {
		if Err := Client.Disconnect(Ctx); Err != nil {
			panic(Err)
		}
	}()
}

func Connect(uri string) {

	Ctx, Cancel = context.WithTimeout(context.Background(), 300*time.Second)

	Client, Err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	// return Client, Ctx, cancel, err
	return
}

func Ping(Client *mongo.Client, Ctx context.Context) error {

	if err := Client.Ping(Ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
