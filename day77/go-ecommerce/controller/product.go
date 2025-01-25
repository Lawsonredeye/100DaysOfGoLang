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

		model.DB.Exec("INSERT INTO products (name, description, sku, category_id, color, product_size, quantity, price, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", product.Name, product.Description, product.SKU, product.CategoryID, product.Color,
			product.ProductSize, product.Quantity, product.Price, product.Created_at, product.Updated_at)
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

// FindProductByName finds all product based on their category name and not product name.
// Using the HTTP GET Method, Products are returned to the client in a List based on the
// items found.
// Parameter:
// - c : *gin.Context
// Response:
// - HTTP 200: List of products.
// - HTTP 404: No product was found.
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

// DeleteProductByID deletes a product from the database along with it's entire
// quantity passed into the database.
// Parameter:
// - c : *gin.Context
// Response:
// - HTTP 204: Product deleted successfully.
// - HTTP 400: Wrong HTTP Method, use DELETE.
func DeleteProductByID(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id := c.Param("id")
	model.DB.Delete(&model.Product{}, id)
	c.JSON(http.StatusNoContent, "")
}

// UpdatedProductByID updated only the field passed by the user for
// updating product information.
// Using the PUT /product/:id, informations about specific products for updates
// are done.
// Parameter:
// - c : *gin.Context
// Response:
// - HTTP 200: Product updated successfully.
// - HTTP 400: Wrong HTTP method, use PUT.
// - HTTP 404: Product with ID not found.
func UpdateProductByID(c *gin.Context) {
	if c.Request.Method != "PUT" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Wrong HTTP method, use PUT.",
		})
		return
	}

	var product model.Product
	var updatedProduct model.Product

	c.BindJSON(&updatedProduct)
	if updatedProduct.Name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result := model.DB.Where("id = ?", c.Param("id")).First(&product)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Product with id [" + c.Param("id") + "] not found.",
		})
		return
	}

	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	if updatedProduct.Price != 0 {
		product.Price = updatedProduct.Price
	}
	if updatedProduct.Color != "" {
		product.Color = updatedProduct.Color
	}
	if updatedProduct.ProductSize != 0 {
		product.ProductSize = updatedProduct.ProductSize
	}
	if updatedProduct.Quantity != 0 {
		product.Quantity = updatedProduct.Quantity
	}
	if updatedProduct.Category != "" {
		var category model.Categories
		err := model.DB.Where("name = ?", updatedProduct.Category).First(&category)
		if err.Error != nil {
			category.Name = "unknown"
			category.ID = 23
		}
		product.SKU = skuGenerator(category.Name, product.Color, product.ProductSize)
		product.CategoryID = int(category.ID)
	}
	product.Updated_at = time.Now()

	model.DB.Model(product).Omit("created_at").Updates(product)

	c.JSON(http.StatusOK, "Product updated successfully")
}
