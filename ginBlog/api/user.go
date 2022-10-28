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
	if err := AddUser2(c); err != nil {
		return
	}
	if err := c.ShouldBind(&user); err != nil {

		response.Failed("参数错误", c)
		return
	}
	fmt.Printf("% + v\n", user.Username)
	fmt.Println(c)
	dao.Mgr.AddUser(&user)

	response.Success("添加成功", user, c)
}

// 创建用户
func AddUser2(con *gin.Context) error {
	var c dao.Manager
	var user models.BlogUser
	//if user.Password != user.Password2 {
	//	return ErrMismatchedPasswords
	//}
	// 用户名存在
	_, err := c.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}
	// 无效用户名
	if models.ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}
	// 无效密码
	if models.ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	// 创建用户
	//c.AddUser(&user)
	return err
}
