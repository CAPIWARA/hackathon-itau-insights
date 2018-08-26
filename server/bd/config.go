package bd

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session
var Database = "capiplusplus"

type MongoCollection struct {
	m *mgo.Collection
}

func ConnectDatabase() error {
	session, err := mgo.Dial("mongodb://capi:Capi10@ds133622.mlab.com:33622/capiplusplus")

	if err != nil {
		log.Printf("db error: %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	Session = session
	return nil
}

func MongoBuilder(tableName string) MongoCollection {
	return MongoCollection{
		m: Session.DB(Database).C(tableName),
	}
}
