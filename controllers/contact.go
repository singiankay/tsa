// contact_controller.go

package Contact

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/singiankay/tsa/models"
	service "github.com/singiankay/tsa/services"
)

type ContactResponse struct {
	ID          int      `json:"id"`
	FullName    string   `json:"full_name" binding:"required"`
	PhoneNumbers []string `json:"phone_numbers"`
	Email       *string  `json:"email,omitempty" binding:"omitempty,email"`
}

var contactService = service.NewContactService()

func CreateContact(c *gin.Context) {
	var input struct {
		FullName    string   `json:"full_name" binding:"required"`
		Email       *string  `json:"email,omitempty" binding:"omitempty,email"`
		PhoneNumbers []string `json:"phone_numbers"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := contactService.CreateContact(input.FullName, input.Email, input.PhoneNumbers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func GetContacts(c *gin.Context) {
	contacts, err := contactService.GetContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []ContactResponse
	for _, contact := range contacts {
		response = append(response, ContactResponse{
			ID:           contact.ID,
			FullName:     contact.FullName,
			PhoneNumbers: flattenPhoneNumbers(contact.PhoneNumbers),
			Email:        contact.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetContactById(c *gin.Context) {
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}
	contactService := service.NewContactService()
	contact, err := contactService.GetContactById(contactID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := ContactResponse{
		ID:           contact.ID,
		FullName:     contact.FullName,
		PhoneNumbers: flattenPhoneNumbers(contact.PhoneNumbers),
		Email:        contact.Email,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func UpdateContact(c *gin.Context) {
	var input struct {
		FullName    string   `json:"full_name" binding:"required"`
		Email       *string  `json:"email,omitempty" binding:"omitempty,email"`
		PhoneNumbers []string `json:"phone_numbers"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}
	contactService := service.NewContactService()
	contact, err := contactService.UpdateContact(contactID, input.FullName, input.Email, input.PhoneNumbers)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func DeleteContact(c *gin.Context) {
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}
	contactService := service.NewContactService()
	if err := contactService.DeleteContact(contactID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func flattenPhoneNumbers(phoneNumberModel []models.PhoneNumber) []string {
	var flatten []string
	for _, phoneNumber := range phoneNumberModel {
		flatten = append(flatten, phoneNumber.Number)
	}
	return flatten
}
