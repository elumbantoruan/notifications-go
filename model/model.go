package model

import (
	"time"
	"github.com/google/uuid"
)

// Represent a Notification structure
type Notification struct {
	Id string
	FirstName string
	LastName string
	Schedule time.Time
	Message string
	PhoneNumber string
	EmailAddress string
	NotificationType string
	Sent bool
}

// New returns an instance of Notification
func (n Notification) New(firstName, lastName string, schedule time.Time, message string, phoneNumber string, emailAddress string, notificationType string) Notification {
	return Notification{
		Id: uuid.New().String(),
		FirstName: firstName,
		LastName: lastName,
		Schedule: schedule,
		Message: message,
		PhoneNumber: phoneNumber,
		EmailAddress: emailAddress,
		NotificationType: notificationType,
		Sent: false,
	}
}
