package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryModel struct {
	ID          string `bson:"id,omitempty" json:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type CategoryModelImpl struct {
	CategoryCollection *mongo.Collection
}

func NewCategoryModelImpl(db *mongo.Database, client *mongo.Client) *CategoryModelImpl {
	return &CategoryModelImpl{
		CategoryCollection: db.Collection("category"),
	}
}

func (c *CategoryModelImpl) GetAllCategory() ([]CategoryModel, error) {
	cursor, err := c.CategoryCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var categories []CategoryModel
	for cursor.Next(context.Background()) {
		var cat CategoryModel
		cursor.Decode(&cat)
		categories = append(categories, cat)
	}
	return categories, nil
}

func (c *CategoryModelImpl) InsertCategory(cat *CategoryModel) (*mongo.InsertOneResult, error) {

	res, err := c.CategoryCollection.InsertOne(context.Background(), cat)

	if err != nil {
		return nil, err
	}

	return res, nil
}
