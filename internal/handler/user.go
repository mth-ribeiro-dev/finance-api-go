package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	User *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{User: userService}
}

// AddUser godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (handler *UserHandler) AddUser(context *gin.Context) {
	var newUser model.User
	if err := context.ShouldBindJSON(&newUser); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		if errors.As(err, &syntaxError) || errors.As(err, &unmarshalTypeErr) {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := handler.User.AddUser(newUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Add users successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// AuthenticateUser godoc
// @Summary Authenticate a user
// @Description Authenticate a user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param loginInfo body object true "Login information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /users/auth [post]
func (handler *UserHandler) AuthenticateUser(context *gin.Context) {
	var loginInfo struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := context.ShouldBindJSON(&loginInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	user, authenticated := handler.User.Authenticate(loginInfo.Email, loginInfo.Password)
	if !authenticated {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete an existing user from the system
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (handler *UserHandler) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = handler.User.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		}
		return
	}

}
