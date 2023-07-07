package controllers

import (
	"gomongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserControllerImpl struct {
	UsermodelImpl *models.UsermodelImpl
}

func NewUserController(db *mongo.Database, client *mongo.Client) *UserControllerImpl {
	UserModel := models.NewUserModelImpl(db, client)
	return &UserControllerImpl{
		UsermodelImpl: UserModel,
	}
}

func (u *UserControllerImpl) InsertUser(ctx *gin.Context) {
	var users models.Usermodel
	ctx.BindJSON(&users)
	users.ID = primitive.NewObjectID().Hex()
	res, err := u.UsermodelImpl.InsertUser(&users)
	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, bson.M{
		"data":    res.InsertedID,
		"message": "success",
	})
}
func (u *UserControllerImpl) GetUser(ctx *gin.Context) {
	var filter = bson.M{}
	res, err := u.UsermodelImpl.FindUsers(filter)

	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, bson.M{
		"data":    res,
		"message": "success",
	})
}
