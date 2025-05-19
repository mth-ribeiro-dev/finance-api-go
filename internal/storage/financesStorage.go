package storage

import (
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
)

type FinanceStorage interface {
	Save([]model.Transaction) error
	Load() ([]model.Transaction, error)
}

type FileFinanceStorage struct {
	FileStorage
}

func NewFileFinanceStorage(filename string) *FileFinanceStorage {
	return &FileFinanceStorage{
		FileStorage: *NewFileStorage(filename),
	}
}

func (f FileFinanceStorage) Save(transactions []model.Transaction) error {
	return f.FileStorage.Save(transactions)
}

func (f FileFinanceStorage) Load() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := f.FileStorage.Load(&transactions)
	return transactions, err
}
