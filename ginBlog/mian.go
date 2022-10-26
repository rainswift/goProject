package main

import (
	"goDemo/ginBlog/dao"
	"goDemo/ginBlog/models"
)

func main() {
	user := models.BlogUser{
		Username: "xyy",
		Password: "123456",
	}
	dao.Mgr.AddUser(&user)

}
