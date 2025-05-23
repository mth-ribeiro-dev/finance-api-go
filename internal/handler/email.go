package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/service"
	"net/http"
)

type EmailHandler struct {
	emailService *service.EmailService
}

func NewEmailHandler(emailService *service.EmailService) *EmailHandler {
	return &EmailHandler{emailService: emailService}
}

// SendEmail godoc
// @Summary Send an email
// @Description Send an email using the provided email data
// @Tags email
// @Accept json
// @Produce json
// @Param emailData body model.EmailData true "Email data"
// @Success 200 {object} map[string]string "message":"Email sent successfully"
// @Failure 400 {object} map[string]string "error":"Invalid email data"
// @Failure 500 {object} map[string]string "error":"Failed to send email"
// @Router /send-email [post]
func (h *EmailHandler) SendEmail(c *gin.Context) {
	var emailData model.EmailData

	if err := c.ShouldBindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email data"})
		return
	}

	err := h.emailService.SendEmail(emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
