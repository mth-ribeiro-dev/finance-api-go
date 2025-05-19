package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/handler"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
	"time"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

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
	router.PUT("/finance/:id", financeHandler.UpdateTransaction)
	router.DELETE("/finance/:id", financeHandler.DeleteTransaction)

	router.POST("/user/register", userHandler.AddUser)
	router.POST("/user/login", userHandler.AuthenticateUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)
}
