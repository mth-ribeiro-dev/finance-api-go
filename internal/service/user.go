package service

import (
	"errors"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/storage"
	"log"
	"strconv"
	"sync"
)

type UserService struct {
	User    []model.User
	NextID  int
	Storage storage.UserStorage
	mu      sync.Mutex
}

func NewUserService(storage storage.UserStorage) *UserService {
	users, err := storage.Load()
	if err != nil {
		log.Printf("Error loading users: %v", err)
		users = []model.User{}
	}

	return &UserService{
		User:    users,
		NextID:  getMaxIDUsers(users) + 1,
		Storage: storage,
	}
}

func (userService *UserService) emailExists(email string) bool {
	for _, user := range userService.User {
		if user.Email == email {
			return true
		}
	}
	return false
}

func getMaxIDUsers(users []model.User) int {
	maxID := 0
	for _, userModel := range users {
		if userModel.ID > maxID {
			maxID = userModel.ID
		}
	}
	return maxID
}

func (userService *UserService) loadUsers() {
	users, err := userService.Storage.Load()
	if err != nil {
		log.Printf("Error loading users file: %v", err)
		userService.User = []model.User{}
		userService.NextID = 1
		return
	}
	userService.User = users
	userService.NextID = getMaxIDUsers(users) + 1
}

func (userService *UserService) saveUser() error {
	err := userService.Storage.Save(userService.User)
	if err != nil {
		log.Printf("Error saving users file: %v", err)
		return err
	}
	return nil
}

func (userService *UserService) AddUser(user model.User) (model.User, error) {
	userService.mu.Lock()
	defer userService.mu.Unlock()
	if userService.emailExists(user.Email) {
		return model.User{}, errors.New("email already exists")
	}
	user.ID = userService.NextID
	user.Status = true
	userService.User = append(userService.User, user)
	userService.NextID++

	err := userService.saveUser()

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userService *UserService) Authenticate(email, password string) (*model.User, bool) {
	userService.mu.Lock()
	defer userService.mu.Unlock()

	for _, userAuth := range userService.User {
		if userAuth.Email == email && userAuth.Password == password && userAuth.Status {
			return &model.User{
				ID:    userAuth.ID,
				Name:  userAuth.Name,
				Email: userAuth.Email,
			}, true
		}
	}
	return nil, false
}

func (userService *UserService) DeleteUser(userId string) error {
	userService.mu.Lock()
	defer userService.mu.Unlock()

	id, _ := strconv.Atoi(userId)
	for index, userModel := range userService.User {
		if userModel.ID == id {
			userModel.Status = false
			userService.User[index] = userModel
			return userService.Storage.Save(userService.User)
		}
	}
	return errors.New("user not found")
}
