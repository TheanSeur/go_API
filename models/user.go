package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Usermodel struct {
	ID       string `json:"_id,omitempty" bson:"_id"`
	UserName string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type UsermodelImpl struct {
	UserCollection *mongo.Collection
}

func NewUserModelImpl(db *mongo.Database, client *mongo.Client) *UsermodelImpl {
	return &UsermodelImpl{
		UserCollection: db.Collection("users"),
	}
}

func (impl *UsermodelImpl) FindUsers(filter bson.M) ([]Usermodel, error) {
	cursor, err := impl.UserCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var users []Usermodel
	for cursor.Next(context.Background()) {
		var user Usermodel
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil

}

func (impl *UsermodelImpl) InsertUser(u *Usermodel) (*mongo.InsertOneResult, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, err
	}
	u.Password = string(hash)

	res, err := impl.UserCollection.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res, nil
}
