package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"mw2020/database"
)

func CustomAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		conn, err := database.GetConnection()
		defer conn.Release()
		if err != nil {
			c.AbortWithStatus(503)
			return
		}

		var userId int
		err = conn.QueryRow(context.Background(), "SELECT user_id FROM sessions WHERE secret=$1",
			c.Request.Header["Authorization"][0]).Scan(&userId)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
