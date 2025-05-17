package service

import (
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"strconv"
	"testing"
	"time"
)

func setupTestService() *FinanceService {
	financeService := NewFinanceService()
	financeService.Transaction = []model.Transaction{}
	financeService.NextID = 1
	return financeService
}

func createTestTransaction() model.Transaction {
	timeDate, _ := time.Parse("2006-01-02", time.Now().Format("yyyy-MM-dd"))
	return model.Transaction{
		Type:        "income",
		Amount:      100.0,
		Category:    "Test",
		Date:        model.DateOnly(timeDate),
		Description: "This is a test transaction",
	}
}

func TestAddTransactionAssignsID(test *testing.T) {
	finance := setupTestService()
	transactionTest := createTestTransaction()
	result := finance.AddTransaction(transactionTest)

	if result.ID != 1 {
		test.Errorf("Expected transaction ID '%v', got '%v'", transactionTest.ID, result.ID)
	}
}

func TestAddTransactionAppendsToList(test *testing.T) {
	finance := setupTestService()
	transactionTest := createTestTransaction()
	finance.AddTransaction(transactionTest)

	if len(finance.Transaction) != 1 {
		test.Errorf("Expected 1 transaction, got %d", len(finance.Transaction))
	}
}

func TestGetAllReturnsTransactions(test *testing.T) {
	finance := setupTestService()
	finance.AddTransaction(createTestTransaction())
	finance.AddTransaction(createTestTransaction())

	result := finance.GetAll()
	if len(result) != 2 {
		test.Errorf("Expected 2 transactions, got %d", len(result))
	}
}

func TestGetBalanceIncomeOnly(test *testing.T) {
	finance := setupTestService()
	transactionTest := createTestTransaction()
	transactionTest.Amount = 150
	finance.AddTransaction(transactionTest)

	if finance.GetBalance() != 150 {
		test.Errorf("Expected 150, got %.2f", finance.GetBalance())
	}
}

func TestGetBalanceIncomeAndExpense(test *testing.T) {
	finance := setupTestService()
	income := createTestTransaction()
	income.Amount = 200
	finance.AddTransaction(income)

	expense := createTestTransaction()
	expense.Type = "expense"
	expense.Amount = 75
	finance.AddTransaction(expense)

	expectedBalance := 125.0
	result := finance.GetBalance()
	if result != expectedBalance {
		test.Errorf("Expected balance %f, got %f", expectedBalance, result)
	}
}

func TestDeleteTransactionSuccess(test *testing.T) {
	finance := setupTestService()
	transactionTest := finance.AddTransaction(createTestTransaction())
	err := finance.DeleteTransaction(strconv.Itoa(transactionTest.ID))
	if err != nil {
		test.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteTransactionNotFound(test *testing.T) {
	finance := setupTestService()
	err := finance.DeleteTransaction("1234")
	if err == nil {
		test.Errorf("Expected 'Transaction not found', got %v", err)
	}
}

func TestUpdateTransactionSuccess(test *testing.T) {
	finance := setupTestService()
	transactionTest := finance.AddTransaction(createTestTransaction())

	updated := createTestTransaction()
	updated.Amount = 700

	err := finance.UpdateTransaction(strconv.Itoa(transactionTest.ID), updated)
	if err != nil {
		test.Errorf("Expected no error, got %v", err)
	}

	if finance.Transaction[0].Amount != updated.Amount {
		test.Errorf("Expected transaction amount %v, got %v", updated.Amount, finance.Transaction[0].Amount)
	}
}

func TestUpdateTransactionNotFound(test *testing.T) {
	finance := setupTestService()
	err := finance.UpdateTransaction("999", model.Transaction{})
	if err == nil {
		test.Errorf("Expected 'Transaction not found', got %v", err)
	}
}
