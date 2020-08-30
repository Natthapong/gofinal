package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Natthapong/gofinal/customer_service/api"
	"github.com/Natthapong/gofinal/customer_service/customer"
	"github.com/Natthapong/gofinal/customer_service/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	database.CreateDatabaseCustomer(db)

	h := customer.Handler{DB: db}
	r := gin.Default()

	r = api.SetupRouter(r, h)
	r.Run(":2009")
}
