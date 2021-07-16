package route

import (
	"net/http"

	"github.com/Z1yx/go-todolist/db"
	"github.com/Z1yx/go-todolist/model"
	v "github.com/Z1yx/go-todolist/variable"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)
func Getall(c *gin.Context) {
	Thing := make([]model.ThingInfo, 0, 10)
	ss := db.Connect()
	defer ss.Close()
	err := ss.DB(v.Db).C(v.Collection).Find(nil).All(&Thing)
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
	Thing := new(model.ThingInfo)
	err := c.Bind(&Thing)
	Thing.ID = bson.NewObjectId()
	ss := db.Connect()
	defer ss.Close()
	err = ss.DB(v.Db).C(v.Collection).Insert(Thing)
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
	ss := db.Connect()
	defer ss.Close()
	err := ss.DB(v.Db).C(v.Collection).Remove(bson.M{"_id": &id})
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

// 此段是壞的 有空一定修
/*func DeleteAll(c *gin.Context) {
	ss := db.Connect()
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