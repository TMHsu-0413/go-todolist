package db

import (
	v "github.com/Z1yx/go-todolist/variable"
	"gopkg.in/mgo.v2"
)

func Connect() *mgo.Session {
	globalS, err := mgo.Dial(v.Host)
	if err != nil {
		panic(err)
	}
	return globalS
}
