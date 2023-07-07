package routes

import (
	"gomongo/controllers"
	"gomongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProductRoutes(route *gin.Engine, db *mongo.Database, client *mongo.Client) {
	productControllerImpl := controllers.ProductControllerImpl{
		ProductModelImpl: models.NewProductModelImpl(db, client),
	}
	r := route.Group("/products")
	r.POST("/all", productControllerImpl.GetProduct)
	r.POST("/add", productControllerImpl.InstertProduct)
	r.PUT("/:id", productControllerImpl.UpdateProduct)
	r.DELETE("/:id", productControllerImpl.DeleteProduct)
}
