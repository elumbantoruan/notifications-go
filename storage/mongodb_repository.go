package storage

import (
	"crypto/tls"
	"log"
	"gopkg.in/mgo.v2"
	"net"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/elumbantoruan/notifications-go/model"
)

type mongodbrepository struct {
	address 			[]string
	database 			string
	username 			string
	password 			string
	databaseCollection 	string
	collectionName		string
	collection 			*mgo.Collection
	session 			*mgo.Session
}


func NewMongoDBRepository(address []string, database, username, password, databaseCollection, collectionName string) (*mongodbrepository, error) {
	r := mongodbrepository{
		address: address,
		database: database,
		username: username,
		password: password,
		databaseCollection:databaseCollection,
		collectionName:collectionName,
	}
	col, err := r.connect()
	if err != nil {
		return nil, err
	}
	r.collection = col

	return &r, nil
}

func (r *mongodbrepository) InsertNotification(notification model.Notification) (interface{}, error) {


	err := r.collection.Insert(notification)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return nil,nil
}

func (r *mongodbrepository) InsertNotifications(notifications []model.Notification) (interface{}, error) {

	docs := make([]interface{}, len(notifications))
	for i,v := range notifications { docs[i] = v}
	err := r.collection.Insert(docs...)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return nil,nil
}

// GetNotifications returns list of notification that hasn't sent (sent:false) and scheduled earlier
func (r *mongodbrepository) GetNotifications() []model.Notification {
	var results []model.Notification

	ft := time.Now()

	err := r.collection.Find(bson.M{
		"sent": false,
		"schedule": bson.M{"$lt": ft}}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results

}

// UpdateNotification (sent:true) for a given id
func (r *mongodbrepository) UpdateNotification(id string, notification *model.Notification) (interface{}, error) {
	selector := bson.M{"id": id}
	change := bson.M{
		"$set": bson.M{
			"id":id,
			"firstname":notification.FirstName,
			"lastname":notification.LastName,
			"schedule":notification.Schedule,
			"message":notification.Message,
			"phonenumber":notification.PhoneNumber,
			"emailaddress":notification.EmailAddress,
			"notificationtype":notification.NotificationType,
			"sent":notification.Sent}}
	e := r.collection.Update(selector, change)
	if e != nil {
		return nil, e
	}
	return nil, nil
}

func (r *mongodbrepository) Close() {
	r.session.Close()
}


func (r *mongodbrepository) connect() (*mgo.Collection, error){

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: r.address,
		Database:r.database,
		Username:r.username,
		Password:r.password,
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	r.session = session

	c := session.DB("notification").C("notifications")

	return c, nil
}



