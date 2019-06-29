package customer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Customer structure
type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func handleError(c *gin.Context, err error, status int) {
	log.Println(err.Error())
	c.JSON(status, gin.H{"error": http.StatusText(status)})
}
