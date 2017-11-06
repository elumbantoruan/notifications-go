package notification

import "github.com/elumbantoruan/notifications-go/model"

type EmailNotification struct {

}

// SendMessage via email
func (e EmailNotification) SendMessage(notification model.Notification) (interface{}, error) {
	return "send notification via email",nil
}

