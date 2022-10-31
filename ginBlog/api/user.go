package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProject/ginBlog/dao"
	"goProject/ginBlog/models"
	response "goProject/ginBlog/responose"
)

type Service struct {
	r dao.Manager
}

func AddUser(c *gin.Context) {

	var user models.BlogUser

	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}
	if err := AddUser2(user.Username); err != nil {
		response.Failed("用户名已经存在", c)
		return
	}
	fmt.Printf("输入的名字 % + v\n", user.Username)
	fmt.Println(c)
	dao.Mgr.AddUser(&user)
	response.Success("添加成功", user, c)
}

// 创建用户
func AddUser2(name string) error {

	// 用户名存在
	flag := dao.Mgr.GetByName(name)
	fmt.Println(flag)
	if flag {
		return nil
	} else {
		return ErrUserExistWithName
	}
	// 无效用户名
	//if models.ValidateUserName(user.Username) {
	//	return ErrInvalidUsername
	//}
	// 无效密码
	//if models.ValidatePassword(user.Password) {
	//	return ErrInvalidPassword
	//}
	// 创建用户
	//c.AddUser(&user)
}
