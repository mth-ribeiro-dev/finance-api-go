package service

import (
	"errors"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

type FinanceService struct {
	Transaction []model.Transaction
	NextID      int
	Filename    string
}

func NewFinanceService() *FinanceService {
	financeService := &FinanceService{}
	financeService.Filename = financeService.getFilePath()
	financeService.loadTransactions()
	return financeService
}

func (financeService *FinanceService) getFilePath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}
	basePath := filepath.Join(currentUser.HomeDir, "financeiro")

	err = os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating financeiro directory: %v", err)
	}
	return filepath.Join(basePath, "transactions.json")
}

func (FinanceService *FinanceService) getMaxID() int {
	maxID := 0
	for _, transaction := range FinanceService.Transaction {
		if transaction.ID > maxID {
			maxID = transaction.ID
		}
	}
	return maxID
}

func (financeService *FinanceService) loadTransactions() {
	transactions, err := storage.LoadFromFile(financeService.Filename)
	if err != nil {
		log.Printf("Error loading finance transactions: %v\n", err)
		financeService.Transaction = []model.Transaction{}
		financeService.NextID = 1
		return
	}
	financeService.Transaction = transactions
	financeService.NextID = financeService.getMaxID() + 1
}

func (financeService *FinanceService) saveTransactions() {
	err := storage.SaveToFile(financeService.Transaction, financeService.Filename)
	if err != nil {
		log.Printf("Error saving finance transactions: %v\n", err)
	}
}

func (finance *FinanceService) AddTransaction(t model.Transaction) model.Transaction {
	t.ID = finance.NextID
	finance.NextID++
	finance.Transaction = append(finance.Transaction, t)
	finance.saveTransactions()
	return t
}

func (finance *FinanceService) GetAll() []model.Transaction {
	return finance.Transaction
}

func (finance *FinanceService) GetBalance() float64 {
	var balance float64
	for _, transaction := range finance.Transaction {
		if transaction.Type == "income" {
			balance += transaction.Amount
		} else if transaction.Type == "expense" {
			balance -= transaction.Amount
		}
	}
	return balance
}

func (finance *FinanceService) DeleteTransaction(idString string) error {
	id, _ := strconv.Atoi(idString)
	for index, transaction := range finance.Transaction {
		if transaction.ID == id {
			finance.Transaction = append(finance.Transaction[:index], finance.Transaction[index+1:]...)
			return storage.SaveToFile(finance.Transaction, finance.Filename)
		}
	}
	return errors.New("Transaction not found")
}

func (finance *FinanceService) UpdateTransaction(idString string, updated model.Transaction) error {
	id, _ := strconv.Atoi(idString)
	for index, transaction := range finance.Transaction {
		if transaction.ID == id {
			updated.ID = id
			finance.Transaction[index] = updated
			return storage.SaveToFile(finance.Transaction, finance.Filename)
		}
	}
	return errors.New("Transaction not found")
}
