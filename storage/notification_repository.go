package storage

import (
	"github.com/elumbantoruan/notifications-go/model"
)

// Represent an interface for notificationRepository
type NotificationRepository interface {
	InsertNotification(notification model.Notification) (interface{}, error)
	InsertNotifications(notifications []model.Notification) (interface{}, error)
	GetNotifications() []model.Notification
	UpdateNotification(id string, notification *model.Notification) (interface{}, error)
}