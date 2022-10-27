package api

import (
	"fmt"
	"net/http"
	"server/dao"
	"server/models"
	"server/responose"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	var user models.UserXyy

	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}
	fmt.Printf("% + v\n", c)
	//fmt.Println(c)
	dao.Mgr.AddUser(&user)

	response.Success("添加成功成功", user, c)
}

func ListUser(c *gin.Context) {
	fmt.Printf("列表数据： % + v\n", c)
	users := dao.Mgr.ListUser()
	response.Success("查询成功", users, c)
}

func GetAllUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.UserXyy
	userLists, err := dao.Mgr.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})

}
