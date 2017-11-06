package factory

import (
	"github.com/elumbantoruan/notifications-go/notification"
	"github.com/elumbantoruan/notifications-go/storage"
	"github.com/google/uuid"
)

type MessagingFactory struct {
	communications map[string]notification.NotificationMessaging
	repository storage.NotificationRepository
}

// NewMessaging returns messaging factory
func NewMessaging(communications map[string]notification.NotificationMessaging, repository storage.NotificationRepository ) MessagingFactory {

	return MessagingFactory{
		communications: communications,
		repository: repository,
	}
}

// SendMessage sends messages
func (m MessagingFactory) SendMessage() (interface{}, error) {

	notifications := m.repository.GetNotifications()

	for _, n := range notifications {
		if _, err := m.communications[n.NotificationType].SendMessage(n); err != nil {
			// TODO: Evaluate the exception.
			// If it's transient then we wouldn't update the notification to true
			// so it will be retry again
			// If not (such as invalid phone number or email address) it should be parked somewhere
		} else {
			id := n.Id
			n.Sent = true
			m.repository.UpdateNotification(id, &n)
		}
	}

	return uuid.New().String(), nil
}
