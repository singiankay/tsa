package service

import (
	"testing"

	"github.com/singiankay/tsa/config"
	"github.com/singiankay/tsa/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	// Use an in-memory SQLite database for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	config.DB = db

	// AutoMigrate models to create the schema
	config.DB.AutoMigrate(&models.Contact{}, &models.PhoneNumber{})
}

func TestCreateContact(t *testing.T) {
	setupTestDB()
	service := NewContactService()

	t.Run("successfully create contact", func(t *testing.T) {
		contact, err := service.CreateContact("John Doe", nil, []string{"+61412345678"})
		assert.NoError(t, err)
		assert.NotNil(t, contact)
		assert.Equal(t, "John Doe", contact.FullName)
		assert.Equal(t, 1, len(contact.PhoneNumbers))
		assert.Equal(t, "+61412345678", contact.PhoneNumbers[0].Number)
	})

	t.Run("fail on invalid phone number", func(t *testing.T) {
		_, err := service.CreateContact("John Doe", nil, []string{"invalid-phone"})
		assert.Error(t, err)
		assert.Equal(t, "invalid phone number format", err.Error())
	})
}

func TestGetContacts(t *testing.T) {
	setupTestDB()
	service := NewContactService()

	// Seed the database
	service.CreateContact("John Doe", nil, []string{"+61412345678"})

	t.Run("successfully retrieve contacts", func(t *testing.T) {
		contacts, err := service.GetContacts()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(contacts))
		assert.Equal(t, "John Doe", contacts[0].FullName)
	})
}

func TestGetContactById(t *testing.T) {
	setupTestDB()
	service := NewContactService()

	// Seed the database
	contact, _ := service.CreateContact("John Doe", nil, []string{"+61412345678"})

	t.Run("successfully retrieve contact by ID", func(t *testing.T) {
		retrievedContact, err := service.GetContactById(contact.ID)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", retrievedContact.FullName)
	})

	t.Run("fail on non-existing contact ID", func(t *testing.T) {
		_, err := service.GetContactById(9999)
		assert.Error(t, err)
		assert.Equal(t, "contact not found", err.Error())
	})
}

func TestUpdateContact(t *testing.T) {
	setupTestDB()
	service := NewContactService()

	// Seed the database
	contact, _ := service.CreateContact("John Doe", nil, []string{"+61412345678"})

	t.Run("successfully update contact", func(t *testing.T) {
		updatedContact, err := service.UpdateContact(contact.ID, "Jane Doe", nil, []string{"+61498765432"})
		assert.NoError(t, err)
		assert.Equal(t, "Jane Doe", updatedContact.FullName)
		assert.Equal(t, 1, len(updatedContact.PhoneNumbers))
		assert.Equal(t, "+61498765432", updatedContact.PhoneNumbers[0].Number)
	})

	t.Run("fail on non-existing contact ID", func(t *testing.T) {
		_, err := service.UpdateContact(9999, "Jane Doe", nil, []string{"+61498765432"})
		assert.Error(t, err)
		assert.Equal(t, "contact not found", err.Error())
	})
}

func TestDeleteContact(t *testing.T) {
	setupTestDB()
	service := NewContactService()

	// Seed the database
	contact, _ := service.CreateContact("John Doe", nil, []string{"+61412345678"})

	t.Run("successfully delete contact", func(t *testing.T) {
		err := service.DeleteContact(contact.ID)
		assert.NoError(t, err)

		// Verify that the contact has been deleted
		_, err = service.GetContactById(contact.ID)
		assert.Error(t, err)
		assert.Equal(t, "contact not found", err.Error())
	})

	t.Run("fail on non-existing contact ID", func(t *testing.T) {
		err := service.DeleteContact(9999)
		assert.Error(t, err)
		assert.Equal(t, "contact not found", err.Error())
	})
}
