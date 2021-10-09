package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type Post struct {
	Id               primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	User             primitive.ObjectID  `json:"user,omitempty" bson:"user,omitempty"`
	Caption          string              `json:"caption,omitempty" bson:"caption,omitempty"`
	Image_URL        string              `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Posted_timestamp primitive.Timestamp `bson:"posted_timestamp,omitempty"`
}
