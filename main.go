package main

import (
	"fmt"
	"net/http"

	"github.com/Natthapong/gofinal/customer"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func auth(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	if authKey != "November 10, 2009" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(auth)
	r.POST("/customers", customer.CreateCustomerHandler)
	r.GET("/customers/:id", customer.FindOneCustomerHandler)
	r.GET("/customers", customer.FindAllCustomerHandler)
	r.PUT("/customers/:id", customer.UpdateCustomerHandler)
	r.DELETE("/customers/:id", customer.DeleteCustomerHandler)
	return r
}

func main() {
	customer.CreateDatabaseCustomer()
	r := setupRouter()
	r.Run(":2009")
	fmt.Println("customer service")
}
