package service

import (
	"errors"
	"strings"
	"testing"

	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
)

type MockStorage struct {
	transactions []model.Transaction
	saveCalled   bool
	loadCalled   bool
	failSave     bool
	failLoad     bool
}

func (m *MockStorage) Save(transactions []model.Transaction) error {
	m.saveCalled = true
	if m.failSave {
		return errors.New("failed to save")
	}
	m.transactions = transactions
	return nil
}

func (m *MockStorage) Load() ([]model.Transaction, error) {
	m.loadCalled = true
	if m.failLoad {
		return nil, errors.New("failed to load")
	}
	return m.transactions, nil
}

func TestAddTransaction(t *testing.T) {
	mockStorage := &MockStorage{transactions: []model.Transaction{}}
	financeService := NewFinanceService(mockStorage)

	initialNextID := financeService.NextID
	initialTransactionCount := len(financeService.Transaction)

	dateStr := "2023-06-15"
	var date model.DateOnly
	err := date.UnmarshalJSON([]byte(`"` + dateStr + `"`))
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	newTransaction := model.Transaction{
		Description: "Test transaction",
		Amount:      100.0,
		Type:        "income",
		Category:    "Test",
		Date:        date,
	}

	result, err := financeService.AddTransaction(newTransaction)

	if err != nil {
		t.Fatalf("Failed to add transaction: %v", err)
	}

	if result.ID != initialNextID {
		t.Errorf("Expected new transaction ID to be %d, got %d", initialNextID, result.ID)
	}

	if financeService.NextID != initialNextID+1 {
		t.Errorf("Expected NextID to increment to %d, got %d", initialNextID+1, financeService.NextID)
	}

	if len(financeService.Transaction) != initialTransactionCount+1 {
		t.Errorf("Expected transaction count to increase by 1, got %d", len(financeService.Transaction))
	}

	if !mockStorage.saveCalled {
		t.Error("Expected saveTransactions to be called")
	}

	lastTransaction := financeService.Transaction[len(financeService.Transaction)-1]
	if lastTransaction.Description != newTransaction.Description ||
		lastTransaction.Amount != newTransaction.Amount ||
		lastTransaction.Type != newTransaction.Type ||
		lastTransaction.Category != newTransaction.Category ||
		lastTransaction.Date != newTransaction.Date {
		t.Error("Added transaction does not match the input")
	}
}

func TestLoadTransactions(t *testing.T) {
	mockTransactions := []model.Transaction{
		{ID: 1, Description: "Test 1", Amount: 100, Type: "income"},
		{ID: 2, Description: "Test 2", Amount: 200, Type: "expense"},
	}

	mockStorage := &MockStorage{transactions: mockTransactions}

	financeService := NewFinanceService(mockStorage)

	financeService.loadTransactions()

	if len(financeService.Transaction) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(financeService.Transaction))
	}

	expectedNextID := 3
	if financeService.NextID != expectedNextID {
		t.Errorf("Expected NextID to be %d, got %d", expectedNextID, financeService.NextID)
	}

	if !mockStorage.loadCalled {
		t.Error("Expected Load() to be called")
	}
}

func TestLoadTransactionsWithFailure(t *testing.T) {
	mockStorage := &MockStorage{failLoad: true}
	financeService := &FinanceService{
		Transaction: []model.Transaction{{ID: 1}},
		NextID:      5,
		Storage:     mockStorage,
	}

	financeService.loadTransactions()

	if len(financeService.Transaction) != 0 {
		t.Errorf("Expected empty transactions, got %d", len(financeService.Transaction))
	}

	if financeService.NextID != 1 {
		t.Errorf("Expected NextID to be reset to 1, got %d", financeService.NextID)
	}

	if !mockStorage.loadCalled {
		t.Error("Load method should have been called")
	}
}

func TestGetTransactionByUserId(t *testing.T) {
	mockTransactions := []model.Transaction{
		{ID: 1, UserID: 1, Description: "Salary", Amount: 3000, Type: "income"},
		{ID: 2, UserID: 1, Description: "Rent", Amount: 1000, Type: "expense"},
		{ID: 3, UserID: 1, Description: "Groceries", Amount: 200, Type: "expense"},
		{ID: 4, UserID: 2, Description: "Bonus", Amount: 500, Type: "income"},
	}

	mockStorage := &MockStorage{transactions: mockTransactions}
	financeService := NewFinanceService(mockStorage)

	result := financeService.GetTransactionByUserId(1)

	if len(result) != 3 {
		t.Errorf("Expected 3 transactions for user ID 1, got %d", len(result))
	}

	expectedTransactions := map[int]string{
		1: "Salary",
		2: "Rent",
		3: "Groceries",
	}

	for _, tx := range result {
		if expectedDesc, ok := expectedTransactions[tx.ID]; ok {
			if tx.Description != expectedDesc {
				t.Errorf("Expected transaction with ID %d to have description '%s', got '%s'", tx.ID, expectedDesc, tx.Description)
			}
			delete(expectedTransactions, tx.ID)
		} else {
			t.Errorf("Unexpected transaction with ID %d found", tx.ID)
		}
	}

	if len(expectedTransactions) > 0 {
		for id, desc := range expectedTransactions {
			t.Errorf("Expected to find transaction with ID %d and description '%s', but it was not present", id, desc)
		}
	}
}

func TestGetBalance(t *testing.T) {
	mockTransactions := []model.Transaction{
		{ID: 1, Description: "Salary", Amount: 3000, Type: "income"},
		{ID: 1, Description: "Rent", Amount: 1000, Type: "expense"},
		{ID: 1, Description: "Freelance", Amount: 500, Type: "income"},
		{ID: 1, Description: "Groceries", Amount: 200, Type: "expense"},
		{ID: 2, Description: "Bonus", Amount: 1000, Type: "income"},
	}

	mockStorage := &MockStorage{transactions: mockTransactions}
	financeService := NewFinanceService(mockStorage)

	balance := financeService.GetBalanceByUserId(1)
	expectedBalance := 2300.0

	if balance != expectedBalance {
		t.Errorf("Expected balance for user ID 1 to be %.2f, got %.2f", expectedBalance, balance)
	}

	balance2 := financeService.GetBalanceByUserId(2)
	expectedBalance2 := 1000.0

	if balance2 != expectedBalance2 {
		t.Errorf("Expected balance for user ID 2 to be %.2f, got %.2f", expectedBalance2, balance2)
	}
}

func TestDeleteTransaction(t *testing.T) {
	mockTransactions := []model.Transaction{
		{ID: 1, Description: "Test 1", Amount: 100, Type: "income"},
		{ID: 2, Description: "Test 2", Amount: 200, Type: "expense"},
		{ID: 3, Description: "Test 3", Amount: 300, Type: "income"},
	}

	mockStorage := &MockStorage{transactions: mockTransactions}

	financeService := NewFinanceService(mockStorage)

	err := financeService.DeleteTransaction("2")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(financeService.Transaction) != 2 {
		t.Errorf("Expected 2 transactions after deletion, got %d", len(financeService.Transaction))
	}

	for _, tx := range financeService.Transaction {
		if tx.ID == 2 {
			t.Error("Transaction with ID 2 should have been deleted")
		}
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save() to be called")
	}

	if len(mockStorage.transactions) != 2 {
		t.Errorf("Expected 2 transactions in storage after deletion, got %d", len(mockStorage.transactions))
	}
}

func TestDeleteNonExistentTransaction(t *testing.T) {
	mockStorage := &MockStorage{transactions: []model.Transaction{
		{ID: 1, Description: "Test 1", Amount: 100, Type: "income"},
		{ID: 2, Description: "Test 2", Amount: 200, Type: "expense"},
	}}
	financeService := NewFinanceService(mockStorage)

	err := financeService.DeleteTransaction("3")

	if err == nil {
		t.Error("Expected an error when deleting a non-existent transaction, got nil")
	} else if err.Error() != "transaction not found" {
		t.Errorf("Expected error message 'transaction not found', got '%s'", err.Error())
	}

	if len(financeService.Transaction) != 2 {
		t.Errorf("Expected 2 transactions after failed deletion, got %d", len(financeService.Transaction))
	}

	if mockStorage.saveCalled {
		t.Error("Save method should not have been called for non-existent transaction")
	}
}

func TestUpdateTransaction(t *testing.T) {
	mockTransactions := []model.Transaction{
		{ID: 1, Description: "Test 1", Amount: 100, Type: "income"},
		{ID: 2, Description: "Test 2", Amount: 200, Type: "expense"},
	}

	dateStr := "2023-06-15"
	var date model.DateOnly
	err := date.UnmarshalJSON([]byte(`"` + dateStr + `"`))
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	mockStorage := &MockStorage{transactions: mockTransactions}

	financeService := NewFinanceService(mockStorage)

	updatedTransaction := model.Transaction{
		Description: "Updated Test 2",
		Amount:      250,
		Type:        "expense",
		Category:    "Updated Category",
		Date:        date,
	}

	err = financeService.UpdateTransaction("2", updatedTransaction)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updatedTx := financeService.Transaction[1]
	if updatedTx.ID != 2 ||
		updatedTx.Description != "Updated Test 2" ||
		updatedTx.Amount != 250 ||
		updatedTx.Type != "expense" ||
		updatedTx.Category != "Updated Category" ||
		updatedTx.Date != date {
		t.Error("Transaction was not updated correctly")
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save() to be called")
	}

	if len(mockStorage.transactions) != 2 {
		t.Errorf("Expected 2 transactions in storage after update, got %d", len(mockStorage.transactions))
	}
}

func TestUpdateNonExistentTransaction(t *testing.T) {
	mockStorage := &MockStorage{transactions: []model.Transaction{
		{ID: 1, Description: "Test 1", Amount: 100, Type: "income"},
		{ID: 2, Description: "Test 2", Amount: 200, Type: "expense"},
	}}
	financeService := NewFinanceService(mockStorage)

	dateStr := "2023-06-15"
	var date model.DateOnly
	err := date.UnmarshalJSON([]byte(`"` + dateStr + `"`))
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	updatedTransaction := model.Transaction{
		Description: "Updated Non-existent",
		Amount:      300,
		Type:        "income",
		Category:    "Test",
		Date:        date,
	}

	err = financeService.UpdateTransaction("3", updatedTransaction)

	if err == nil {
		t.Error("Expected an error when updating a non-existent transaction, got nil")
	} else if err.Error() != "transaction not found" {
		t.Errorf("Expected error message 'transaction not found', got '%s'", err.Error())
	}

	if len(financeService.Transaction) != 2 {
		t.Errorf("Expected 2 transactions after failed update, got %d", len(financeService.Transaction))
	}

	if mockStorage.saveCalled {
		t.Error("Save method should not have been called for non-existent transaction")
	}
}

func TestSaveTransactionsWithFailure(t *testing.T) {
	mockStorage := &MockStorage{failSave: true}
	financeService := &FinanceService{
		Transaction: []model.Transaction{
			{ID: 1, Description: "Test", Amount: 100, Type: "income"},
		},
		Storage: mockStorage,
	}

	err := financeService.saveTransactions()

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if err != nil && !strings.Contains(err.Error(), "Error saving finance transactions") {
		t.Errorf("Expected error message to contain 'Error saving finance transactions', got: %v", err)
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save method to be called")
	}
}
