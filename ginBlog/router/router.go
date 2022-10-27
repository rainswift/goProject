package router

import (
	"github.com/gin-gonic/gin"
	"goProject/ginBlog/controller"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/users", controller.ListUser)
	e.POST("/users", controller.AddUser)
	e.Run()
}
