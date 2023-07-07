package controllers

import (
	"fmt"
	"gomongo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// controller
type ProductControllerImpl struct {
	ProductModelImpl *models.ProductModelImpl
}

func NewProductController(db *mongo.Database, client *mongo.Client) *ProductControllerImpl {
	productModel := models.NewProductModelImpl(db, client)

	return &ProductControllerImpl{
		ProductModelImpl: productModel,
	}
}

func (p *ProductControllerImpl) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.ProductModel
	res, err := p.ProductModelImpl.DeleteProduct(id, &product)
	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, bson.M{
		"data":    res.DeletedCount,
		"message": "success",
	})
}

func (p *ProductControllerImpl) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.ProductModel
	ctx.BindJSON(&product)
	res, err := p.ProductModelImpl.UpdateProduct(id, &product)
	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, bson.M{
		"data":    res.ModifiedCount,
		"message": "success",
	})
}

func (p *ProductControllerImpl) InstertProduct(ctx *gin.Context) {
	var product models.ProductModel
	ctx.BindJSON(&product)
	product.ID = primitive.NewObjectID().Hex()
	res, err := p.ProductModelImpl.InstertProduct(&product)
	if err != nil {
		ctx.JSON(400, bson.M{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bson.M{
		"mesasge": "success",
		"data":    res.InsertedID,
	})
}

type ProductFilter struct {
	Price float64 `json:"price"`
}

func (p *ProductControllerImpl) GetProduct(ctx *gin.Context) {
	//get all products from mongodb with criteria (filter)
	//all produts return from mongodb is a cursor ([element])
	//delcare variable to retrieve that
	//price
	var filter = bson.M{}
	var productFilter = ProductFilter{
		Price: 0,
	}

	params := ctx.Param("filterPrice")

	ctx.BindJSON(&productFilter)

	if productFilter.Price > 0 {
		filter["price"] = bson.M{
			"$gte": params,
		}
	}
	fmt.Printf("%v", filter)
	result, err := p.ProductModelImpl.FindProducts(filter)
	// pro, err := p.ProductModelImpl.FindProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, bson.M{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, bson.M{
		"data":    result,
		"message": "seccse",
	})
}
