package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"mw2020/database"
	"mw2020/requests"
)

//PostTransaction creates new transaction between users
func PostTransaction(c *gin.Context) {
	var err error
	var ctx = context.Background()
	conn, err := database.GetConnection()
	defer conn.Release()
	if err != nil {
		c.AbortWithStatus(503)
		return
	}

	currentUserId, _ := c.Get("user_id")

	var requestBody requests.PostTransactionRequest
	err = c.Bind(&requestBody)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var recipientId int
	err = conn.QueryRow(ctx, "select id from users where email=$1", requestBody.UserEmail).Scan(&recipientId)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, "select * from users where id=$1 for update;", currentUserId)
	_, err = tx.Exec(ctx, "select * from users where id=$1 for update;", recipientId)

	_, err = tx.Exec(ctx, "update users set balance=balance-$1 where id=$2;", requestBody.TransactionAmount, currentUserId)
	_, err = tx.Exec(ctx, "update users set balance=balance+$1 where id=$2", requestBody.TransactionAmount, recipientId)

	_, err = tx.Exec(ctx, "insert into transactions (user_id_from, user_id_to, amount) values ($1, $2, $3);", currentUserId, recipientId, requestBody.TransactionAmount)
	err = tx.Commit(ctx)

	if err != nil {
		_ = tx.Rollback(ctx)
		log.Println("Can`t make the transaction")
		c.AbortWithStatus(503)
		return
	}

	c.Status(200)
}
