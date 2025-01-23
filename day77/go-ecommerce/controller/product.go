package controller

import (
	"fmt"
	"go-commerce/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	// get the product from into an obj
	var product model.Product
	c.BindJSON(&product)
	if product.Category == "" || product.Color == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Category or Color field can not be empty.",
		})
		return
	}
	
	model.DB.Where("name = ?", strings.ToLower(product.Category)).First(&product.CategoryID)
	// pass the data into the skuGenerator for a unique ID

	sku := skuGenerator(product.Category, product.Color, product.ProductSize)
	// store the sku into the obj.SKU
	product.SKU = sku
	// store the obj into the db
	rows := model.DB.Create(&product).RowsAffected

	if rows <= 0 {
		c.Abort()
		return
	}
	// return a 201 response to client
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added successful.",
		"product_details": fmt.Sprintf("%+v", product),
	})
}

func skuGenerator(category string, color string, size int) string {
	// category JEANS, color BLUE, 12 => JNS-BLU-12
	ccode, ok := ColorCode[strings.ToLower(color)]
	if !ok {
		ccode = "UNK"
	}

	categoryCode, ok := CategoryAbbreviation[strings.ToLower(category)]
	if !ok {
		categoryCode = "UNK"
	}
	return fmt.Sprintf("%v-%v-%04d", categoryCode, ccode, size)
}
