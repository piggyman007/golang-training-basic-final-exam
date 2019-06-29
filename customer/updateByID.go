package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/golang-training-basic-final-exam/database"
)

// UpdateByID - update customer by ID
func UpdateByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	customer := new(Customer)

	if err := c.ShouldBindJSON(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	customer.ID = id

	err = db.UpdateCustomerByID(customer.ID, customer.Name, customer.Email, customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, customer)
}
