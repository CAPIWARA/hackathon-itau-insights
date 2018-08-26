package bd

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session
var Database = "capiplusplus"

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
