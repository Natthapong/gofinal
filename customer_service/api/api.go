package api

import (
	"github.com/Natthapong/gofinal/customer_service/customer"
	"github.com/Natthapong/gofinal/customer_service/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, h customer.Handler) *gin.Engine {
	r.Use(middleware.Auth)
	r.POST("/customers", h.CreateCustomerHandler)
	r.GET("/customers/:id", h.FindOneCustomerHandler)
	r.GET("/customers", h.FindAllCustomerHandler)
	r.PUT("/customers/:id", h.UpdateCustomerHandler)
	r.DELETE("/customers/:id", h.DeleteCustomerHandler)
	return r
}
