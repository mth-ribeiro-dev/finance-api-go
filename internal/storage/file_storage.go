package storage

import (
	"encoding/json"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"os"
)

func SaveToFile(Transaction []model.Transaction, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(Transaction)
}

func LoadFromFile(filename string) ([]model.Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Transaction{}, nil
		}
		return nil, err
	}
	defer file.Close()
	var transactions []model.Transaction
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&transactions)
	return transactions, err
}
