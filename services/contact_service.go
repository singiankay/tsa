package service

import (
	"errors"

	"github.com/singiankay/tsa/config"
	"github.com/singiankay/tsa/models"
	validator "github.com/singiankay/tsa/validators"
)

type ContactService struct{}

func NewContactService() *ContactService {
	return &ContactService{}
}

func (s *ContactService) CreateContact(fullName string, email *string, phoneNumbers []string) (*models.Contact, error) {
	var contact models.Contact
	contact.FullName = fullName
	if email != nil {
		contact.Email = email
	}

	for _, phoneNumber := range phoneNumbers {
		libPhoneValidator := validator.NewPhoneNumberValidator(true)
		number, err := libPhoneValidator.Validate(phoneNumber)
		if err != nil {
			return nil, errors.New("invalid phone number format")
		}
		contact.PhoneNumbers = append(contact.PhoneNumbers, models.PhoneNumber{Number: number})
	}

	if err := config.DB.Create(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (s *ContactService) GetContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	if err := config.DB.Preload("PhoneNumbers").Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *ContactService) GetContactById(id int) (*models.Contact, error) {
	var contact models.Contact
	if err := config.DB.Preload("PhoneNumbers").Where("id = ?", id).First(&contact).Error; err != nil {
		return nil, errors.New("contact not found")
	}
	return &contact, nil
}

func (s *ContactService) UpdateContact(id int, fullName string, email *string, phoneNumbers []string) (*models.Contact, error) {
	var contact models.Contact

	if err := config.DB.Where("id = ?", id).First(&contact).Error; err != nil {
		return nil, errors.New("contact not found")
	}

	// Remove old phone numbers
	if err := config.DB.Where("contact_id =?", contact.ID).Delete(&models.PhoneNumber{}).Error; err != nil {
		return nil, errors.New("error updating records")
	}

	contact.FullName = fullName
	if email != nil {
		contact.Email = email
	}

	for _, phoneNumber := range phoneNumbers {
		libPhoneValidator := validator.NewPhoneNumberValidator(true)
		number, err := libPhoneValidator.Validate(phoneNumber)
		if err != nil {
			return nil, errors.New("invalid phone number format")
		}
		contact.PhoneNumbers = append(contact.PhoneNumbers, models.PhoneNumber{Number: number})
	}

	if err := config.DB.Model(&contact).Updates(contact).Error; err != nil {
		return nil, errors.New("error updating records")
	}

	return &contact, nil
}

func (s *ContactService) DeleteContact(id int) error {
	var contact models.Contact
	if err := config.DB.Preload("PhoneNumbers").Where("id = ?", id).First(&contact).Error; err != nil {
		return errors.New("contact not found")
	}

	if err := config.DB.Select("PhoneNumbers").Delete(&contact).Error; err != nil {
		return errors.New("error deleting contact")
	}

	return nil
}
