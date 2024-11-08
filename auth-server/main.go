package main

import (
	"github.com/Superm4n97/oms/auth-server/database"
	"github.com/Superm4n97/oms/auth-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	r := gin.Default()
	r.GET("/auth/users", handlers.GetUsers)
	r.POST("/auth/users", handlers.CreateUser)

	r.Run(":8080")
}
