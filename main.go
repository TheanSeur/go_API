package main

import (
	"context"
	"gomongo/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductModel struct {
	ID    string  `json:"_id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

func main() {
	r := gin.Default()
	//allow origin
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	client, err := loadMongodbConfig()
	if err != nil {
		panic(err.Error())
	}
	//connect to mongodb client
	db := client.Database("test")
	//creata a gateway for get
	//get all products from mongodb with criteria (filter)
	//all produts return from mongodb is a cursor ([element])
	//delcare variable to retrieve that
	routes.ProductRoutes(r, db, client)
	routes.UserRoutes(r, db, client)
	r.Run(":6969")

}

func loadMongodbConfig() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
