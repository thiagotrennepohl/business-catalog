package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string
	AddressZip int
	Website    string
}
