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

	router.Run(":8081")

}
