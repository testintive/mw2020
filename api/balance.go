package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"mw2020/database"
	"mw2020/responses"
	"net/http"
)

//GetBalance returns info about user including current balance
func GetBalance(c *gin.Context) {
	var err error
	conn, err := database.GetConnection()
	defer conn.Release()
	if err != nil {
		c.AbortWithStatus(503)
		return
	}

	var fetchedBalance int
	userIdFromMiddleware, _ := c.Get("user_id")
	err = conn.QueryRow(context.Background(),
		"SELECT balance FROM users WHERE id=$1", userIdFromMiddleware).Scan(&fetchedBalance)
	if err != nil {
		c.AbortWithStatus(503)
		return
	}

	c.JSON(http.StatusOK, responses.GetBalanceResponse{Balance: fetchedBalance})

}
