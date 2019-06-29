package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/golang-training-basic-final-exam/database"
)

// GetByID - get customer by ID
func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	row, err := db.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	customer := new(Customer)
	if err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, customer)
}
