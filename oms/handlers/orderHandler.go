package handlers

import (
	"context"
	"github.com/Superm4n97/oms/oms/database"
	"github.com/Superm4n97/oms/oms/models"
	"github.com/Superm4n97/oms/order-processor/handlers"
	"github.com/Superm4n97/oms/order-processor/queue"
	"k8s.io/klog/v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
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

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO orders (username, description) VALUES ($1, $2)", order.Username, order.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err = queue.PushToQueue(handlers.OrderQueue, order); err != nil {
		klog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Order created"})
}
