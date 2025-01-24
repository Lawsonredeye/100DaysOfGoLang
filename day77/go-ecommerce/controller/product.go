package controller

import (
	"fmt"
	"go-commerce/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AddProduct adds new product into the database and also adds
// a SKU for identification of products.
// Listens for a POST request to handle with a JSON body for key values
// such as product [name, description, cagtegory, color, size, quantity].
// Parameter:
// - c : *gin.Context - GIN framework
// Response:
// - HTTP 201: Product added successfully
// - HTTP 400: Empty required fields.
func AddProduct(c *gin.Context) {
	if c.Request.Method == "POST" {
		// get the product from into an obj
		var product model.Product
		c.BindJSON(&product)
		if product.Category == "" || product.Color == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Category or Color field can not be empty.",
			})
			return
		}

		var abvr model.Categories

		model.DB.Where("name = ?", strings.ToLower(product.Category)).First(&abvr)

		if abvr.ID == 0 {
			product.CategoryID = 23
		} else {
			product.CategoryID = int(abvr.ID)
		}

		fmt.Println(product.CategoryID)

		// pass the data into the skuGenerator for a unique ID
		sku := skuGenerator(product.Category, product.Color, product.ProductSize)
		product.SKU = sku

		product.Created_at = time.Now()
		product.Updated_at = time.Now()

		model.DB.Exec("INSERT INTO products (name, description, sku, category_id, color, product_size, quantity, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ? ,?)", product.Name, product.Description, product.SKU, product.CategoryID, product.Color,
			product.ProductSize, product.Quantity, product.Created_at, product.Updated_at)
		// return a 201 response to client
		c.JSON(http.StatusCreated, gin.H{
			"message":         "Product added successful.",
			"product_details": product,
		})
	}
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

// GET 200 /product/:name
func FindProductByName(c *gin.Context) {
	if c.Request.Method == "GET" {
		// using the endpoint search for all products that has the id in the products database
		// return 200
		name := c.Param("name")

		var categoryName model.Categories

		model.DB.Where("name = ?", name).First(&categoryName)

		if categoryName.Name != "" {
			var allProducts []model.Product

			model.DB.Where("category_id = ?", categoryName.ID).Find(&allProducts)
			c.JSON(http.StatusOK, allProducts)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Oops! Nothing Found.",
		})
		return
	}
}
