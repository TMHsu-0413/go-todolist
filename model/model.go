package model

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "Todolist"
	collection = "Todolist"
	host       = "127.0.0.1:27017"
)

var globalS *mgo.Session

type ThingInfo struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"Name"`
}

func init() {
	globalS, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	fmt.Println(globalS)
}
