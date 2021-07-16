package main

import (
	"github.com/Z1yx/go-todolist/middleware"
	"github.com/Z1yx/go-todolist/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	v1 := router.Group("/api")
	{
		v1.GET("/", route.Getall)
		v1.POST("/", route.Create)
		v1.DELETE("/:id", route.Delete)
		//v1.PUT("/:id", route.Update)
		//v1.DELETE("All", DeleteAll)
	}
	router.Run(":8888")
}
