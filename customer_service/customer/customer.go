package customer

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Natthapong/gofinal/customer_service/database"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

type Handler struct {
	DB *sql.DB
}

func (h *Handler) CreateCustomerHandler(c *gin.Context) {
	cust := database.Customer{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id := database.InsertCustomer(h.DB, &cust)
	cust.ID = id
	c.JSON(http.StatusCreated, cust)
}

func (h *Handler) FindOneCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	cust := database.FindCustomerByID(h.DB, id)
	c.JSON(http.StatusOK, &cust)
}

func (h *Handler) FindAllCustomerHandler(c *gin.Context) {
	cust := database.FindCustomers(h.DB)
	c.JSON(http.StatusOK, &cust)
}

func (h *Handler) UpdateCustomerHandler(c *gin.Context) {
	cust := database.Customer{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	cust = database.UpdateCustomer(h.DB, cust.ID, cust.Name, cust.Email, cust.Status)
	c.JSON(http.StatusOK, &cust)
}

func (h *Handler) DeleteCustomerHandler(c *gin.Context) {
	msg := Message{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	errDelete := database.DeleteCustomer(h.DB, id)
	if errDelete == nil {
		msg.Message = "customer deleted"
	}
	c.JSON(http.StatusOK, msg)
}
