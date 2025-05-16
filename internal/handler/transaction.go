package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"net/http"
)

type Handler struct {
	Finance *service.FinanceService
}

func NewHandler(finance *service.FinanceService) *Handler {
	return &Handler{Finance: finance}
}

func (handler *Handler) AddTransaction(context *gin.Context) {
	var transaction model.Transaction
	if err := context.ShouldBindJSON(&transaction); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Data"})
		return
	}
	if transaction.Type != "income" && transaction.Type != "expense" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Type must be income or expense"})
		return
	}
	handler.Finance.AddTransaction(transaction)
	context.JSON(http.StatusCreated, transaction)
}

func (handler *Handler) GetTransactions(context *gin.Context) {
	context.JSON(http.StatusOK, handler.Finance.GetAll())
}

func (handler *Handler) GetBalance(context *gin.Context) {
	context.JSON(http.StatusOK, handler.Finance.GetBalance())
}
