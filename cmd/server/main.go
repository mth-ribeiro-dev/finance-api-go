package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/handler"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
)

func main() {
	router := gin.Default()

	setupServices(router)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

func setupServices(router *gin.Engine) {

	financeStorage := storage.NewFileFinanceStorage("finances.json")
	financeService := service.NewFinanceService(financeStorage)
	financeHandler := handler.NewFinanceHandler(financeService)

	userStorage := storage.NewFileUserStorage("users.json")
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/finance/transaction", financeHandler.AddTransaction)
	router.GET("/finance/transactions/:userId", financeHandler.GetTransactions)
	router.GET("/finance/balance/:userId", financeHandler.GetBalance)
	router.PUT("/transactions/:id", financeHandler.UpdateTransaction)
	router.DELETE("/transactions/:id", financeHandler.DeleteTransaction)

	router.POST("/user/register", userHandler.AddUser)
	router.POST("/user/login", userHandler.AuthenticateUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)
}
