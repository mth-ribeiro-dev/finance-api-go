package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type FinanceHandle struct {
	Finance *service.FinanceService
}

func NewFinanceHandler(finance *service.FinanceService) *FinanceHandle {
	return &FinanceHandle{Finance: finance}
}

// AddTransaction godoc
// @Summary Add a new transaction
// @Description Add a new financial transaction for a user
// @Tags finance
// @Accept json
// @Produce json
// @Param transaction body model.Transaction true "Transaction object"
// @Success 201 {object} model.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
func (handler *FinanceHandle) AddTransaction(context *gin.Context) {
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

	transaction, err := handler.Finance.AddTransaction(transaction)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add transaction"})
		return
	}
	context.JSON(http.StatusCreated, transaction)
}

// GetTransactions godoc
// @Summary Get user transactions
// @Description Get all transactions for a specific user
// @Tags finance
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} model.Transaction
// @Failure 400 {object} map[string]string
// @Router /transactions/{userId} [get]
func (handler *FinanceHandle) GetTransactions(context *gin.Context) {
	userIDStr := context.Param("userId")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	transactions := handler.Finance.GetTransactionByUserId(userID)
	context.JSON(http.StatusOK, transactions)
}

// GetBalance godoc
// @Summary Get user balance
// @Description Get the current balance for a specific user
// @Tags finance
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Router /balance/{userId} [get]
func (handler *FinanceHandle) GetBalance(context *gin.Context) {
	userIDStr := context.Param("userId")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	balance := handler.Finance.GetBalanceByUserId(userID)
	context.JSON(http.StatusOK, gin.H{"balance": balance})
}

// UpdateTransaction godoc
// @Summary Update a transaction
// @Description Update an existing financial transaction
// @Tags finance
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param transaction body model.Transaction true "Updated transaction object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [put]
func (handler *FinanceHandle) UpdateTransaction(context *gin.Context) {
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
		if err.Error() == "transaction not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete an existing financial transaction
// @Tags finance
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [delete]
func (handler *FinanceHandle) DeleteTransaction(context *gin.Context) {
	id := context.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = handler.Finance.DeleteTransaction(id)
	if err != nil {
		if err.Error() == "transaction not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
