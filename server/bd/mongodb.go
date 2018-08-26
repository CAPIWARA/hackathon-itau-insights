package bd

import (
	"errors"

	"github.com/cardosomarcos/structs"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (db *MongoCollection) FindById(id string) (interface{}, error) {
	var data map[string]interface{}

	if err := db.m.Find(bson.M{"_id": id}).One(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (db *MongoCollection) Save(obj interface{}) (string, error) {
	payload := structs.New(obj).Map()

	id := bson.NewObjectId()
	payload["_id"] = id.Hex()

	if err := db.m.Insert(payload); err != nil {
		if mgo.IsDup(err) {
			return "", errors.New("record already exists!")
		}
		return "", err
	}

	return id.Hex(), nil
}
