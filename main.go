package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Natthapong/gofinal/customer"
	"github.com/Natthapong/gofinal/database"
	"github.com/Natthapong/gofinal/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter(r *gin.Engine, h customer.Handler) *gin.Engine {
	r.Use(middleware.Auth)
	r.POST("/customers", h.CreateCustomerHandler)
	r.GET("/customers/:id", h.FindOneCustomerHandler)
	r.GET("/customers", h.FindAllCustomerHandler)
	r.PUT("/customers/:id", h.UpdateCustomerHandler)
	r.DELETE("/customers/:id", h.DeleteCustomerHandler)
	return r
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	database.CreateDatabaseCustomer(db)

	h := customer.Handler{DB: db}
	r := gin.Default()

	r = setupRouter(r, h)
	r.Run(":2009")
}
