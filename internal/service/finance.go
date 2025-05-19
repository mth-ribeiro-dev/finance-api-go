package service

import (
	"errors"
	"fmt"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
	"log"
	"strconv"
	"sync"
)

type FinanceService struct {
	Transaction []model.Transaction
	NextID      int
	Storage     storage.FinanceStorage
	mu          sync.Mutex
}

func NewFinanceService(storage storage.FinanceStorage) *FinanceService {
	transactions, err := storage.Load()
	if err != nil {
		log.Printf("Error loading transactions: %v\n", err)
		transactions = []model.Transaction{}
	}

	return &FinanceService{
		Transaction: transactions,
		NextID:      getMaxIDTransactions(transactions) + 1,
		Storage:     storage,
	}
}

func getMaxIDTransactions(transactions []model.Transaction) int {
	maxID := 0
	for _, transactionModel := range transactions {
		if transactionModel.ID > maxID {
			maxID = transactionModel.ID
		}
	}
	return maxID
}

func (financeService *FinanceService) loadTransactions() {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	transactions, err := financeService.Storage.Load()
	if err != nil {
		log.Printf("Error loading finance transactions: %v\n", err)
		financeService.Transaction = []model.Transaction{}
		financeService.NextID = 1
		return
	}
	financeService.Transaction = transactions
	financeService.NextID = getMaxIDTransactions(transactions) + 1
}

func (financeService *FinanceService) saveTransactions() error {
	err := financeService.Storage.Save(financeService.Transaction)
	if err != nil {
		return fmt.Errorf("Error saving finance transactions: %v\n", err)
	}
	return nil
}

func (financeService *FinanceService) AddTransaction(transaction model.Transaction) (model.Transaction, error) {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	transaction.ID = financeService.NextID
	financeService.NextID++
	financeService.Transaction = append(financeService.Transaction, transaction)
	err := financeService.saveTransactions()
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (financeService *FinanceService) GetTransactionByUserId(userID int) []model.Transaction {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	var result []model.Transaction
	for _, transaction := range financeService.Transaction {
		if transaction.UserID == userID {
			result = append(result, transaction)
		}
	}
	return result
}

func (financeService *FinanceService) GetBalanceByUserId(userID int) float64 {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	var balance float64
	for _, transaction := range financeService.Transaction {
		if transaction.ID == userID {
			if transaction.Type == "income" {
				balance += transaction.Amount
			} else if transaction.Type == "expense" {
				balance -= transaction.Amount
			}
		}
	}
	return balance
}

func (financeService *FinanceService) DeleteTransaction(idString string) error {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	id, _ := strconv.Atoi(idString)
	for index, transactionModel := range financeService.Transaction {
		if transactionModel.ID == id {
			financeService.Transaction = append(financeService.Transaction[:index], financeService.Transaction[index+1:]...)
			return financeService.Storage.Save(financeService.Transaction)
		}
	}
	return errors.New("transaction not found")
}

func (financeService *FinanceService) UpdateTransaction(idString string, updated model.Transaction) error {
	financeService.mu.Lock()
	defer financeService.mu.Unlock()

	id, _ := strconv.Atoi(idString)
	for index, transactionModel := range financeService.Transaction {
		if transactionModel.ID == id {
			updated.ID = id
			financeService.Transaction[index] = updated
			return financeService.Storage.Save(financeService.Transaction)
		}
	}
	return errors.New("transaction not found")
}
