package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storable interface {
	Save(data interface{}) error
	Load(data interface{}) error
}

type FileStorage struct {
	Filename string
}

func NewFileStorage(filename string) *FileStorage {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	appDir := filepath.Join(homeDir, "myfinance")
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		panic(err)
	}
	return &FileStorage{
		Filename: filepath.Join(appDir, filename),
	}
}

func (f FileStorage) Save(data interface{}) error {
	file, err := os.Create(f.Filename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	return err
}

func (f FileStorage) Load(data interface{}) error {
	file, err := os.Open(f.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	return err
}
