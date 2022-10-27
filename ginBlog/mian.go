package main

import (
	"goProject/ginBlog/dao"
	"goProject/ginBlog/models"
)

func main() {
	user := models.BlogUser{
		Username: "xyy",
		Password: "123456",
	}
	dao.Mgr.AddUser(&user)
}
