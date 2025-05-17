package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"net/http"
	"strings"
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
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		if errors.As(err, &syntaxErr) || errors.As(err, &unmarshalTypeErr) {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON syntax"})
			return
		}

		if strings.Contains(err.Error(), "date") {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Date must be in the format yyyy-mm-dd"})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if transaction.Type != "income" && transaction.Type != "expense" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Type must be 'income' or 'expense'"})
		return
	}

	transaction = handler.Finance.AddTransaction(transaction)
	context.JSON(http.StatusCreated, transaction)
}

func (handler *Handler) GetTransactions(context *gin.Context) {
	context.JSON(http.StatusOK, handler.Finance.GetAll())
}

func (handler *Handler) GetBalance(context *gin.Context) {
	context.JSON(http.StatusOK, handler.Finance.GetBalance())
}

func (handler *Handler) UpdateTransaction(context *gin.Context) {
	id := context.Param("id")
	var updatedTransaction model.Transaction
	if err := context.ShouldBindJSON(&updatedTransaction); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		if errors.As(err, &syntaxErr) || errors.As(err, &unmarshalTypeErr) {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON syntax"})
			return
		}

		if strings.Contains(err.Error(), "date") {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Date must be in the format yyyy-mm-dd"})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if updatedTransaction.Type != "income" && updatedTransaction.Type != "expense" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Type must be 'income' or 'expense'"})
		return
	}

	err := handler.Finance.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

func (handler *Handler) DeleteTransaction(context *gin.Context) {
	id := context.Param("id")

	err := handler.Finance.DeleteTransaction(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
