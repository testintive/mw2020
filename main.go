package main

import (
	"github.com/gin-gonic/gin"
	"mw2020/api"
	"mw2020/database"
	"mw2020/middleware"
)

func main() {
	var err error
	err = database.InitDb()
	if err != nil {
		panic("Can`t init DB")
	}

	r := gin.Default()
	r.Use(middleware.CustomAuth())

	securedApi := r.Group("/api")
	{
		securedApi.GET("/balance", api.GetBalance)
		securedApi.POST("/transaction", api.PostTransaction)
	}

	err = r.Run()
	if err != nil {
		panic("Can`t start server")
	}
}
