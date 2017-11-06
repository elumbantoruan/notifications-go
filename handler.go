package main

import (
	"log"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/elumbantoruan/notifications-go/notification"
	"encoding/json"
	"os"
	"github.com/elumbantoruan/notifications-go/storage"
	"github.com/elumbantoruan/notifications-go/factory"
	"github.com/elumbantoruan/notifications-go/model"
	"time"
)

func main() {
	// Lambda will not execute main function
	// This is only needed for development purpose

	// The following environment variables need to be added in AWS lambda environment variables setup
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")

	os.Setenv("MONGODB_CLUSTER0", "")
	os.Setenv("MONGODB_CLUSTER1", "")
	os.Setenv("MONGODB_CLUSTER2", "")
	os.Setenv("MONGODB_DATABASE", "")
	os.Setenv("MONGODB_USERNAME", "")
	os.Setenv("MONGODB_PASSWORD", "")
	os.Setenv("MONGODB_DATABASECOLLECTION", "")
	os.Setenv("MONGODB_COLLECTIONNAME", "")


	r, e := Handler(nil, nil)

	if e != nil {
		log.Println(e.Error())
	} else {
		log.Println(r)
	}
}

// Handler is the entry point for AWS Lambda
func Handler(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Println("start")

	address := []string {
		os.Getenv("MONGODB_CLUSTER0"),
		os.Getenv("MONGODB_CLUSTER1"),
		os.Getenv("MONGODB_CLUSTER2"),
	}
	database := os.Getenv("MONGODB_DATABASE")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	databaseCollection := os.Getenv("MONGODB_DATABASECOLLECTION")
	collectionName := os.Getenv("MONGODB_COLLECTIONNAME")

	nm := notification.NewNotificationMessaging()
	nr, err := storage.NewMongoDBRepository(address, database, username, password, databaseCollection, collectionName)
	if err != nil {
		return nil, err
	}

	nf := factory.NewMessaging(nm, nr)
	r, e := nf.SendMessage()
	if e != nil {
		return nil, e
	}
	return r,nil

}

func insertNotification() error {

	address := []string {
		os.Getenv("MONGODB_CLUSTER0"),
		os.Getenv("MONGODB_CLUSTER1"),
		os.Getenv("MONGODB_CLUSTER2"),
	}
	database := os.Getenv("MONGODB_DATABASE")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	databaseCollection := os.Getenv("MONGODB_DATABASECOLLECTION")
	collectionName := os.Getenv("MONGODB_COLLECTIONNAME")

	n := model.Notification{}
	notifications := []model.Notification{
		n.New("John", "Doe", time.Now().AddDate(0,0,1), "test", "1 999 999 9999", "", "sms"),
		n.New("Nancy", "Doe", time.Now().AddDate(0,0,1), "test2", "1 999 999 9999", "", "sms"),
	}

	var err error

	nr, err := storage.NewMongoDBRepository(address, database, username, password, databaseCollection, collectionName)

	if err != nil {
		return err
	}

	_, err = nr.InsertNotifications(notifications)
	if err != nil {
		return err
	}
	return nil
}
