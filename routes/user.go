package routes

import (
	"gomongo/controllers"
	"gomongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(routes *gin.Engine, db *mongo.Database, client *mongo.Client) {
	userControllerImpl := controllers.UserControllerImpl{
		UsermodelImpl: models.NewUserModelImpl(db, client),
	}
	r := routes.Group("/users")
	r.GET("/", userControllerImpl.GetUser)
	r.POST("/add", userControllerImpl.InsertUser)
}
