package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mth-ribeiro-dev/finance-api-go.git/docs"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/config"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/handler"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	SMTP struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
}

func loadConfig(filename string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// @title MyFinance API
// @version 0.3.3
// @description This is a REST API for managing personal finances developed in Go.

// @contact.name Matheus Ribeiro
// @contact.email matheus.junio159@gmail.com

// @license.name Creative Commons BY-NC 4.0
// @license.url https://creativecommons.org/licenses/by-nc/4.0/

// @host localhost:8081
// @BasePath /api/v1
// @schemes http
func main() {

	router := gin.Default()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:8080"}
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	configCors.AllowCredentials = true
	configCors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configCors))

	setupServices(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

// setupServices configures and sets up all the services for the API
func setupServices(router *gin.Engine) {

	// Finance service setup
	financeStorage := storage.NewFileFinanceStorage("finances.json")
	financeService := service.NewFinanceService(financeStorage)
	financeHandler := handler.NewFinanceHandler(financeService)

	// User service setup
	userStorage := storage.NewFileUserStorage("users.json")
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	cfg := config.GetConfig()

	// Configurar o serviço de email
	emailService := service.NewEmailService(
		cfg.SMTP.Host,
		cfg.SMTP.Port,
		cfg.SMTP.Username,
		cfg.SMTP.Password,
	)

	emailHandler := handler.NewEmailHandler(emailService)

	v1 := router.Group("/api/v1")
	{
		// Finance routes
		v1.POST("/transactions", financeHandler.AddTransaction)
		v1.GET("/transactions/:userId", financeHandler.GetTransactions)
		v1.GET("/balance/:userId", financeHandler.GetBalance)
		v1.PUT("/transactions/:id", financeHandler.UpdateTransaction)
		v1.DELETE("/transactions/:id", financeHandler.DeleteTransaction)

		// User routes
		v1.POST("/users", userHandler.AddUser)
		v1.POST("/users/auth", userHandler.AuthenticateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)

		// Email routes
		v1.POST("/send-email", emailHandler.SendEmail)
	}
}
