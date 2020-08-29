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

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

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

func CreateDatabaseCustomer() {
	createTb := `
		CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY, 
		name TEXT,
		email TEXT,
		status TEXT
		);
	`
	_, err := db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
}

func insertCustomer(cust *Customer) (id int) {
	row := db.QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3) RETURNING id", cust.Name, cust.Email, cust.Status)
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("Error can't scan id", err)
	}
	return
}

func findCustomerByID(id int) (cust Customer) {
	cust = Customer{}
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		log.Fatal("Error can't prepare query one row statement", err)
	}
	row := stmt.QueryRow(id)
	err = row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
	if err != nil {
		log.Println("Error can't scan row into variables >>", err)
	}
	return
}

func findCustomers() (customers []Customer) {
	customers = []Customer{}
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		log.Fatal("Error can't prepare query all rows statement", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error can't query all rows", err)
	}
	for rows.Next() {
		cust := Customer{}
		err := rows.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
		if err != nil {
			log.Fatal("Error can't scan row into variables", err)
		}
		customers = append(customers, cust)
	}
	return
}

func updateCustomer(id int, name, email, status string) (cust Customer) {
	stmt, err := db.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1")
	if err != nil {
		log.Fatal("Error can't prepare statment update", err)
	}
	if _, err := stmt.Exec(id, name, email, status); err != nil {
		log.Println("Error execute update ", err)
	}
	cust = findCustomerByID(id)
	return cust
}

func deleteCustomer(id int) (err error) {
	stmt, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		log.Fatal("Error can't prepare statment delete", err)
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Println("Error error execute delete ", err)
	}
	return
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
