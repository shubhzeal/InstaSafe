package main

import (
	"github/InstaSafe/Controllers"
	"github/InstaSafe/Model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Model.ConnectDatabase()

	r.POST("/transactions", Controllers.CreateTransaction)

	r.Run()
}
