package main

import (
	"fmt"

	"github.com/Natthapong/gofinal/customer"
	"github.com/Natthapong/gofinal/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Auth)
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
