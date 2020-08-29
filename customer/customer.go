package customer

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Message struct {
	Message string `json:"message"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	//defer db.Close()
}

func CreateCustomerHandler(c *gin.Context) {
	cust := Customer{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id := insertCustomer(&cust)
	cust.ID = id
	c.JSON(http.StatusCreated, cust)
}

func FindOneCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	cust := findCustomerByID(id)
	c.JSON(http.StatusOK, &cust)
}

func FindAllCustomerHandler(c *gin.Context) {
	cust := findCustomers()
	c.JSON(http.StatusOK, &cust)
}

func UpdateCustomerHandler(c *gin.Context) {
	cust := Customer{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	cust = updateCustomer(cust.ID, cust.Name, cust.Email, cust.Status)
	c.JSON(http.StatusOK, &cust)
}

func DeleteCustomerHandler(c *gin.Context) {
	msg := Message{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	errDelete := deleteCustomer(id)
	if errDelete == nil {
		msg.Message = "customer deleted"
	}
	c.JSON(http.StatusOK, msg)
}
