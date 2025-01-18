package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var colorAbvr = make(map[string]string, 0)

func main() {
	fmt.Println()

	router := gin.Default()

	router.GET("/", createProduct)
	router.Run(":3030")
}

type ProductDetail struct {
	ID          int
	Name        string
	Quantity    int
	Category    string
	Color       string
	Description string
	Size        int
	SKU         string
	Price       float64
}

func createProduct(c *gin.Context) {
	var product ProductDetail
	c.BindJSON(&product)

	if product.Category == "" || product.Color == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Category or Color type is missing",
		})
		return
	}
	fmt.Println(product)
	product.SKU = skuGenerator(product.Category, product.Color, product.Size)
	c.JSON(200, product)
}

func skuGenerator(c string, co string, s int) string {
	colorAbvr["red"] = "RED"
	colorAbvr["green"] = "GRN"
	colorAbvr["blue"] = "BLU"
	colorAbvr["pink"] = "PNK"
	colorAbvr["white"] = "WHT"

	c = abvrGen(c)
	color := colorAbvr[co]
	return fmt.Sprintf("%v-%v-%v", c, color, s)
}

func abvrGen(s string) string {
	strArr := strings.Split(strings.ToUpper(s), "")
	str := fmt.Sprintf("%v%v%v", strArr[0], strArr[1], strArr[2])

	return str
}
