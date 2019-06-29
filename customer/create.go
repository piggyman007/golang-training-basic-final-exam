package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/golang-training-basic-final-exam/database"
)

// Create customer handler
func Create(c *gin.Context) {
	customer := new(Customer)

	if err := c.ShouldBindJSON(customer); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	row, err := db.CreateCustomer(customer.Name, customer.Email, customer.Status)

	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	if err = row.Scan(&customer.ID); err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, customer)
}
