package storage

import (
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
)

type UserStorage interface {
	Save([]model.User) error
	Load() ([]model.User, error)
}

type FileUserStorage struct {
	FileStorage
}

func NewFileUserStorage(filename string) *FileUserStorage {
	return &FileUserStorage{
		FileStorage: *NewFileStorage(filename),
	}
}

func (f FileUserStorage) Save(users []model.User) error {
	return f.FileStorage.Save(users)
}

func (f FileUserStorage) Load() ([]model.User, error) {
	var users []model.User
	err := f.FileStorage.Load(&users)
	return users, err
}
