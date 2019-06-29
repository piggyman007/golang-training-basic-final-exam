package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/piggyman007/golang-training-basic-final-exam/customer"
	db "github.com/piggyman007/golang-training-basic-final-exam/database"
)

func main() {
	db.New()
	defer db.Close()
	r := setupRouter()
	r.Run(":2019")
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware)
	r.POST("/customers", customer.Create)
	r.GET("/customers", customer.List)
	r.GET("/customers/:id", customer.GetByID)
	r.PUT("/customers/:id", customer.UpdateByID)
	r.DELETE("/customers/:id", customer.DeleteByID)

	return r
}
