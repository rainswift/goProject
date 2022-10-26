package controller

import (
	"github.com/gin-gonic/gin"
	"goDemo/ginBlog/dao"
	"goDemo/ginBlog/models"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.BlogUser{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "user.html", nil)
}
