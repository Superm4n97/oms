package handlers

import (
	"context"
	"github.com/Superm4n97/oms/auth-server/database"
	"github.com/Superm4n97/oms/auth-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := database.DB.Exec(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2)", user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "User created"})
}
