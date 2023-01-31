package Controllers

import (
	"github/InstaSafe/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	//validate input
	var input CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//create Transaction
	transaction := Model.Transaction{Amount: input.Amount, Timestamp: input.Timestamp}
	Model.DB.create(&transaction)
	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
