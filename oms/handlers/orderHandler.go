package handlers

import (
	"context"
	"github.com/Superm4n97/oms/oms/database"
	"github.com/Superm4n97/oms/oms/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, username, description FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Username, &order.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		orders = append(orders, order)
	}
	c.JSON(http.StatusOK, orders)
}

func CreateProduct(c *gin.Context) {
	var user models.Order
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO users (username, description) VALUES ($1, $2)", user.Username, user.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "User created"})
}
