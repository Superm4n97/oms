package main

import (
	"github.com/Superm4n97/oms/oms/database"
	"github.com/Superm4n97/oms/oms/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	r := gin.Default()
	r.GET("/product/order", handlers.GetProduct)
	r.POST("/product/order", handlers.CreateProduct)

	r.Run(":9090")
}
