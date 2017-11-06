package notification

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/elumbantoruan/notifications-go/model"
)

type TextNotification struct {}

// SendMessage sends SMS via SNS
func (t TextNotification) SendMessage(notification model.Notification) (interface{}, error) {

	message := notification.Message
	phoneNumber := notification.PhoneNumber

	log.Println(message)
	log.Println(phoneNumber)

	// the credential for client session is pulled from environment variable
	// credentials.NewEnvCredentials() will look for the followings:
	// AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
	clientSession := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
					Credentials:credentials.NewEnvCredentials(),
				},
			}))

	svc := sns.New(clientSession)
	input := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message: aws.String(message),
	}
	po, err := svc.Publish(input)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	log.Println(po.String())
	return po.String(), nil
}