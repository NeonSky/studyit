package util

import (
	"fmt"
	"github.com/globalsign/mgo"
)

const SERVER = "mongodb://database:27017"
const DBNAME = "studyit"

func OpenSession() *mgo.Session {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to server.", err)
	}
	return session
}

func FetchCollection(collectionName string) (*mgo.Collection, *mgo.Session) {
	session := OpenSession()
	c := session.DB(DBNAME).C(collectionName)
	return c, session
}
