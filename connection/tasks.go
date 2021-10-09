package connection

import (
	"log"

	"github.com/chirayurathi/task-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(docs models.User) (models.User, error) {
	user := models.User{}

	collection := Client.Database(DataBase).Collection("user")
	result, err := collection.InsertOne(Ctx, docs)

	if err != nil {
		return user, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return GetUser(id)
}

func GetUser(id string) (models.User, error) {
	user := models.User{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	collection := Client.Database(DataBase).Collection("user")
	err = collection.FindOne(Ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertPost(docs models.Post) (models.Post, error) {
	post := models.Post{}

	// pid, err := primitive.ObjectIDFromHex(docs.User)
	// if err != nil {
	// 	return post, err
	// }
	// docs.User = pid
	collection := Client.Database(DataBase).Collection("post")
	result, err := collection.InsertOne(Ctx, docs)

	if err != nil {
		return post, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return GetPost(id)
}

func GetPost(id string) (models.Post, error) {
	post := models.Post{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return post, err
	}

	collection := Client.Database(DataBase).Collection("post")
	err = collection.FindOne(Ctx, bson.M{"_id": _id}).Decode(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func GetAllPost(id string) ([]models.Post, error) {
	posts := []models.Post{}

	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return posts, err
	}

	collection := Client.Database(DataBase).Collection("post")
	cur, err := collection.Find(Ctx, bson.M{"user": pid})
	if err != nil {
		return posts, err
	}
	defer cur.Close(Ctx)

	for cur.Next(Ctx) {

		var post models.Post
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return posts, nil
}
