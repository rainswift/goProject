package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/models"
	"strings"
)

type Manager interface {
	AddUser(user *models.UserXyy)
	ListUser() []models.UserXyy
	GetAllUsers(user *models.UserXyy, pagination *models.Pagination) (*[]models.UserXyy, error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.UserXyy{})
}

// 添加用户
func (mgr *manager) AddUser(user *models.UserXyy) {
	fmt.Printf("% + v\n", user)
	mgr.db.Create(user)
}

// 查找用户
func (mgr *manager) ListUser() []models.UserXyy {
	var users []models.UserXyy
	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	mgr.db.Limit(10).Offset(10).Find(&users)
	return users
}

//func whereFn(pagination *models.Pagination) (query interface{}, args interface{}) {
//	query = "name like ? AND gender like ? AND age >= ? && age <= ?"
//	args =
//}

func strArr(str string) []string {
	var desStr1 []string
	if str != "" {
		desStr1 = strings.Split(str, ",")
		if len(desStr1) == 1 {
			desStr1 = append(desStr1, "9999")
		}
	} else {
		desStr1 = []string{"0", "9999"}
	}
	//fmt.Printf("% + v\n", desStr1)
	return desStr1
}

func (mgr *manager) GetAllUsers(user *models.UserXyy, pagination *models.Pagination) (*[]models.UserXyy, error) {
	var users []models.UserXyy
	// 分页查询
	offset := (pagination.Page - 1) * pagination.Limit

	queryBuider := mgr.db.Limit(pagination.Limit).Offset(offset)

	result := queryBuider.Model(&models.UserXyy{}).Where(
		"name like ? AND gender like ? AND age >= ? && age <= ? AND height >= ? && height <= ? ",
		"%"+pagination.Sort+"%",
		pagination.Gender+"%",
		strArr(pagination.Age)[0], strArr(pagination.Age)[1],
		strArr(pagination.Height)[0], strArr(pagination.Height)[1],
	).Find(&users)

	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	affected := result.RowsAffected
	fmt.Printf("% + v\n", affected)
	return &users, nil
}
