package controller

import (
	"fmt"
	"go-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// AddOrder takes the product id and the quantity of the product the user
// wants to place.
// It then responds with the price of the total quantity being purchased.
// Taking a post request, it checks if the user is a registered user before
// processing the order.
// Parameter:
// - c : *gin.Context
// Response:
// - HTTP 200: Product added to cart successfully.
// - HTTP 401: User is not authorized to carry out the process
func AddOrder(c *gin.Context) {
	cookie, err:= c.Cookie(CookieToken)
	if err != nil {
		abortWithStatusJSON("Cookie", c)
		return
	}

	userID, err := ValidateCookies(cookie)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "You do not have access")
		return
	}

	var product model.Product
	var orderItem model.OrderItem
	var obj Obj
	
	c.MustBindWith(&obj, binding.Default("POST", "application/json"))

	model.DB.Raw("SELECT * FROM products WHERE id = ?", obj.ProductID).Scan(&product)

	orderItem.Price = product.Price * float64(obj.Quantity)
	orderItem.Quantity = obj.Quantity
	orderItem.ProductID = product.ID

	
	var orders model.Orders
	
	orders.Status = "pending"
	orders.TotalAmount = orderItem.Price
	orders.UserID = uint(userID)

	count := 0
	model.DB.Raw("SELECT COUNT(id) FROM orders").Scan(&count)
	orderItem.OrderID = count + 1
	
	model.DB.Create(&orders)
	model.DB.Create(&orderItem)

	fmt.Println(orderItem)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success. Item added to cart",
		"details": orderItem,
	})
}