package main

import (
	"github.com/Superm4n97/oms/oms/database"
	"github.com/Superm4n97/oms/oms/handlers"
	"github.com/Superm4n97/oms/order-processor/queue"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()
	queue.SetRedisClient()

	r := gin.Default()
	r.GET("/product/orders", handlers.GetOrders)
	r.POST("/product/order", handlers.CreateOrder)

	r.Run(":9090")
}
