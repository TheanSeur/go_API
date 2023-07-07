package controllers

import (
	"gomongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryControllerImpl struct {
	CategoryModelImpl *models.CategoryModelImpl
}

func NewCategoryController(db *mongo.Database, client *mongo.Client) *CategoryControllerImpl {
	category := models.NewCategoryModelImpl(db, client)

	return &CategoryControllerImpl{
		CategoryModelImpl: category,
	}
}

func (c *CategoryControllerImpl) InsertCategory(ctx *gin.Context) {

	var category models.CategoryModel
	ctx.BindJSON(&category)
	category.ID = primitive.NewObjectID().Hex()
	res, err := c.CategoryModelImpl.InsertCategory(&category)

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

func (c *CategoryControllerImpl) GetAllCategory(ctx *gin.Context) {
	res, err := c.CategoryModelImpl.GetAllCategory()

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
