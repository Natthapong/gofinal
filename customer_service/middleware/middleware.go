package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	if authKey != "November 10, 2009" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
}
