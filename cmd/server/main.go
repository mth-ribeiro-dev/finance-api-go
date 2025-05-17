package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/handler"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
)

func main() {
	router := gin.Default()
	financeService := service.NewFinanceService()
	hand := handler.NewHandler(financeService)

	router.POST("/finance/transaction", hand.AddTransaction)
	router.GET("finance/transactions", hand.GetTransactions)
	router.GET("finance/balance", hand.GetBalance)
	router.PUT("/transactions/:id", hand.UpdateTransaction)
	router.DELETE("/transactions/:id", hand.DeleteTransaction)

	err := router.Run(":8081")
	if err != nil {
		return
	}

}
