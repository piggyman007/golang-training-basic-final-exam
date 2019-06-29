package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/golang-training-basic-final-exam/database"
)

// List customer handler
func List(c *gin.Context) {
	rows, err := db.ListCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		customer := new(Customer)
		if err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}

		customers = append(customers, *customer)
	}

	c.JSON(http.StatusOK, customers)
}
