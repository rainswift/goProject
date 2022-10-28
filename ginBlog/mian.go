package main

import "goProject/ginBlog/router"

func main() {
	//user := models.BlogUser{
	//	Username: "xyy",
	//	Password: "123456",
	//}
	//dao.Mgr.AddUser(&user)
	router.Start()
}
