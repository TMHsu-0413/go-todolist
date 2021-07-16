package model

import "gopkg.in/mgo.v2/bson"

type ThingInfo struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"Name"`
}