package notification

import (
	"github.com/elumbantoruan/notifications-go/model"
)

// NotificationMessaging represents an interface for sending message
type NotificationMessaging interface {
	SendMessage(notification model.Notification) (interface{}, error)
}

// NewNotificationMessaging returns lists of notification messaging (sms or email)
func NewNotificationMessaging() map[string]NotificationMessaging {
	return map[string]NotificationMessaging {
		"sms": TextNotification{},
		"email": EmailNotification{},
	}
}
