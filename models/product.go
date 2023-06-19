package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductModel struct {
	ID    string  `bson:"_id,omitempty" json:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

type CreateProductModel struct {
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

type ProductModelImpl struct {
	ProductCollection *mongo.Collection
}

func NewProductModelImpl(db *mongo.Database, client *mongo.Client) *ProductModelImpl {
	return &ProductModelImpl{
		ProductCollection: db.Collection("product"),
	}
}

func (p *ProductModelImpl) UpdateProduct(id string, pro *ProductModel) (*mongo.UpdateResult, error) {
	res, err := p.ProductCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": &pro})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *ProductModelImpl) DeleteProduct(id string, pro *ProductModel) (*mongo.DeleteResult, error) {
	res, err := p.ProductCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *ProductModelImpl) FindProducts() ([]ProductModel, error) {

	cursor, err := p.ProductCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	//declare variable to store all products
	var products []ProductModel
	//loop through cursor and append each element to products
	for cursor.Next(context.Background()) {
		var product ProductModel
		cursor.Decode(&product)
		products = append(products, product)
	}
	//return products
	return products, nil
}

func (p *ProductModelImpl) InstertProduct(pro *ProductModel) (*mongo.InsertOneResult, error) {

	res, err := p.ProductCollection.InsertOne(context.Background(), pro)
	if err != nil {
		return nil, err
	}

	return res, nil
}
