package service

import (
	"errors"
	"testing"

	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
)

type MockUserStorage struct {
	users      []model.User
	saveCalled bool
	loadCalled bool
	failSave   bool
	failLoad   bool
}

func (m *MockUserStorage) Save(users []model.User) error {
	m.saveCalled = true
	if m.failSave {
		return errors.New("failed to save")
	}
	m.users = users
	return nil
}

func (m *MockUserStorage) Load() ([]model.User, error) {
	m.loadCalled = true
	if m.failLoad {
		return nil, errors.New("failed to load")
	}
	return m.users, nil
}

func TestAddUser(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{}}
	userService := NewUserService(mockStorage)

	newUser := model.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	addedUser, err := userService.AddUser(newUser)

	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	if addedUser.ID != 1 {
		t.Errorf("Expected new user ID to be 1, got %d", addedUser.ID)
	}

	if addedUser.Name != newUser.Name {
		t.Errorf("Expected user name to be %s, got %s", newUser.Name, addedUser.Name)
	}

	if addedUser.Email != newUser.Email {
		t.Errorf("Expected user email to be %s, got %s", newUser.Email, addedUser.Email)
	}

	if !addedUser.Status {
		t.Error("Expected user status to be true")
	}

	if len(userService.User) != 1 {
		t.Errorf("Expected user count to be 1, got %d", len(userService.User))
	}

	if userService.NextID != 2 {
		t.Errorf("Expected NextID to be 2, got %d", userService.NextID)
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save() to be called")
	}
}

func TestAddUserWithExistingEmail(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "Existing User", Email: "existing@example.com", Password: "password123", Status: true},
	}}
	userService := NewUserService(mockStorage)

	newUser := model.User{
		Name:     "New User",
		Email:    "existing@example.com",
		Password: "newpassword123",
	}

	_, err := userService.AddUser(newUser)

	if err == nil {
		t.Error("Expected an error when adding a user with an existing email, but got nil")
	}

	if err != nil && err.Error() != "email already exists" {
		t.Errorf("Expected error message 'email already exists', got '%s'", err.Error())
	}

	if len(userService.User) != 1 {
		t.Errorf("Expected user count to remain 1, got %d", len(userService.User))
	}

	if mockStorage.saveCalled {
		t.Error("Save method should not have been called for existing email")
	}
}

func TestAuthenticateUser(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123", Status: true},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password456", Status: false},
	}}
	userService := NewUserService(mockStorage)

	// Test successful authentication
	authenticatedUser, isAuthenticated := userService.Authenticate("john@example.com", "password123")

	if !isAuthenticated {
		t.Error("Expected user to be authenticated, but authentication failed")
	}

	if authenticatedUser == nil {
		t.Fatal("Expected authenticated user to be non-nil, but got nil")
	}

	if authenticatedUser.ID != 1 {
		t.Errorf("Expected authenticated user ID to be 1, got %d", authenticatedUser.ID)
	}

	if authenticatedUser.Name != "John Doe" {
		t.Errorf("Expected authenticated user name to be 'John Doe', got '%s'", authenticatedUser.Name)
	}

	if authenticatedUser.Email != "john@example.com" {
		t.Errorf("Expected authenticated user email to be 'john@example.com', got '%s'", authenticatedUser.Email)
	}

	_, isAuthenticated = userService.Authenticate("jane@example.com", "password456")

	if isAuthenticated {
		t.Error("Expected authentication to fail for inactive user, but it succeeded")
	}
}

func TestAuthenticateWithIncorrectCredentials(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "correctpassword", Status: true},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password456", Status: false},
	}}
	userService := NewUserService(mockStorage)

	_, isAuthenticated := userService.Authenticate("john@example.com", "incorrectpassword")

	if isAuthenticated {
		t.Error("Expected authentication to fail with incorrect password, but it succeeded")
	}

	_, isAuthenticated = userService.Authenticate("nonexistent@example.com", "anypassword")

	if isAuthenticated {
		t.Error("Expected authentication to fail with non-existent email, but it succeeded")
	}

	_, isAuthenticated = userService.Authenticate("jane@example.com", "password456")

	if isAuthenticated {
		t.Error("Expected authentication to fail for inactive user, but it succeeded")
	}
}

func TestAuthenticateInactiveUser(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123", Status: false},
	}}
	userService := NewUserService(mockStorage)

	_, isAuthenticated := userService.Authenticate("john@example.com", "password123")

	if isAuthenticated {
		t.Error("Expected authentication to fail for inactive user, but it succeeded")
	}
}

func TestDeleteUser(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123", Status: true},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "password456", Status: true},
	}}
	userService := NewUserService(mockStorage)

	err := userService.DeleteUser("1")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(userService.User) != 2 {
		t.Errorf("Expected 2 users after deletion, got %d", len(userService.User))
	}

	deletedUser := userService.User[0]
	if deletedUser.ID != 1 || deletedUser.Status != false {
		t.Errorf("Expected user with ID 1 to be marked as inactive, got ID: %d, Status: %v", deletedUser.ID, deletedUser.Status)
	}

	activeUser := userService.User[1]
	if activeUser.ID != 2 || activeUser.Status != true {
		t.Errorf("Expected user with ID 2 to remain active, got ID: %d, Status: %v", activeUser.ID, activeUser.Status)
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save() to be called")
	}
}

func TestDeleteNonExistentUser(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123", Status: true},
	}}
	userService := NewUserService(mockStorage)

	err := userService.DeleteUser("2")

	if err == nil {
		t.Error("Expected an error when deleting a non-existent user, got nil")
	} else if err.Error() != "user not found" {
		t.Errorf("Expected error message 'user not found', got '%s'", err.Error())
	}

	if len(userService.User) != 1 {
		t.Errorf("Expected user count to remain 1, got %d", len(userService.User))
	}

	if mockStorage.saveCalled {
		t.Error("Save method should not have been called for non-existent user")
	}
}

func TestNewUserServiceLoadsUsers(t *testing.T) {
	mockUsers := []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123", Status: true},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "password456", Status: true},
	}

	mockStorage := &MockUserStorage{users: mockUsers}
	userService := NewUserService(mockStorage)

	if len(userService.User) != 2 {
		t.Errorf("Expected 2 users to be loaded, got %d", len(userService.User))
	}

	if userService.NextID != 3 {
		t.Errorf("Expected NextID to be 3, got %d", userService.NextID)
	}

	if !mockStorage.loadCalled {
		t.Error("Expected Load() to be called during service initialization")
	}

	for i, user := range mockUsers {
		if userService.User[i].ID != user.ID || userService.User[i].Name != user.Name || userService.User[i].Email != user.Email {
			t.Errorf("User at index %d does not match expected user", i)
		}
	}
}

func TestNewUserServiceHandlesLoadError(t *testing.T) {
	mockStorage := &MockUserStorage{failLoad: true}
	userService := NewUserService(mockStorage)

	if len(userService.User) != 0 {
		t.Errorf("Expected empty user list when loading fails, got %d users", len(userService.User))
	}

	if userService.NextID != 1 {
		t.Errorf("Expected NextID to be 1 when loading fails, got %d", userService.NextID)
	}

	if !mockStorage.loadCalled {
		t.Error("Expected Load() to be called during service initialization")
	}
}

func TestAddUserSavesToStorage(t *testing.T) {
	mockStorage := &MockUserStorage{users: []model.User{}}
	userService := NewUserService(mockStorage)

	newUser := model.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	_, err := userService.AddUser(newUser)

	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	if !mockStorage.saveCalled {
		t.Error("Expected Save() to be called after adding a new user")
	}

	if len(mockStorage.users) != 1 {
		t.Errorf("Expected 1 user to be saved in storage, got %d", len(mockStorage.users))
	}

	savedUser := mockStorage.users[0]
	if savedUser.Name != newUser.Name || savedUser.Email != newUser.Email || savedUser.Password != newUser.Password {
		t.Error("Saved user does not match the added user")
	}

	if savedUser.ID != 1 {
		t.Errorf("Expected saved user ID to be 1, got %d", savedUser.ID)
	}

	if !savedUser.Status {
		t.Error("Expected saved user status to be true")
	}
}

func TestNewUserServiceDeterminesNextID(t *testing.T) {
	mockUsers := []model.User{
		{ID: 1, Name: "User 1", Email: "user1@example.com"},
		{ID: 3, Name: "User 3", Email: "user3@example.com"},
		{ID: 2, Name: "User 2", Email: "user2@example.com"},
	}

	mockStorage := &MockUserStorage{users: mockUsers}
	userService := NewUserService(mockStorage)

	expectedNextID := 4
	if userService.NextID != expectedNextID {
		t.Errorf("Expected NextID to be %d, got %d", expectedNextID, userService.NextID)
	}

	if !mockStorage.loadCalled {
		t.Error("Expected Load() to be called during service initialization")
	}

	if len(userService.User) != len(mockUsers) {
		t.Errorf("Expected %d users to be loaded, got %d", len(mockUsers), len(userService.User))
	}
}
