package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err == nil {
		fmt.Println("connected!")
	} else {
		fmt.Println(err)
	}
	session.SetPoolLimit(10)
}
