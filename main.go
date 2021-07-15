package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "Todolist"
	collection = "Todolist"
	host       = "127.0.0.1:27017"
)

type ThingInfo struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"Name"`
}

func connect() *mgo.Session {
	globalS, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	return globalS
}

func main() {
	server := gin.Default()
	server.GET("/", Getall)
	server.POST("/", Create)
	server.Run(":8888")
}

func Getall(c *gin.Context) {
	Thing := make([]ThingInfo, 0, 10)
	ss := connect()
	defer ss.Close()
	err := ss.DB(db).C(collection).Find(nil).All(&Thing)
	fmt.Println(Thing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Not exist",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": &Thing,
	})
}

func Create(c *gin.Context) {
	Thing := new(ThingInfo)
	err := c.Bind(&Thing)
	Thing.ID = bson.NewObjectId()
	ss := connect()
	defer ss.Close()
	err = ss.DB(db).C(collection).Insert(Thing)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid request Body",
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Accepted",
		"message": "Add data",
	})
}
