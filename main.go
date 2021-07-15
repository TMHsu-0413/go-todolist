package main

import (
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
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/", Getall)
		v1.POST("/", Create)
		v1.DELETE("/:id", Delete)
		v1.PUT("/:id", Update)
		//v1.DELETE("All", DeleteAll)
	}
	router.Run(":8888")
}

func Getall(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	Thing := make([]ThingInfo, 0, 10)
	ss := connect()
	defer ss.Close()
	err := ss.DB(db).C(collection).Find(nil).All(&Thing)
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

func Delete(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	ss := connect()
	defer ss.Close()
	err := ss.DB(db).C(collection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error in the Thing id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Delete completed!",
	})
}
func Update(c *gin.Context) {

}

// 此段是壞的 有空一定修
/*func DeleteAll(c *gin.Context) {
	ss := connect()
	defer ss.Close()
	err := ss.DB(db).C(collection).Remove(bson.M{"_id": nil})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error in the Thing id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Delete All completed!",
	})
}*/
