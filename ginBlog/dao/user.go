package dao

import (
	"goProject/ginBlog/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.BlogUser)
	GetByName(name string) (models.BlogUser, error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.BlogUser{})
}

//func (mgr *manager) AddUser(user *models.BlogUser) {
//	fmt.Printf("% + v\n", user)
//	mgr.db.Create(user)
//}

// 创建用户
func (mgr *manager) AddUser(user *models.BlogUser) {
	mgr.db.Create(user)

}

// 根据用户名查询用户
func (mgr *manager) GetByName(name string) (models.BlogUser, error) {
	var user models.BlogUser
	err := mgr.db.Where("UserName = ?", name).Where("IsDeleted = ?", 0).First(&user, "UserName = ?", name).Error
	if err != nil {
		return models.BlogUser{}, err
	}
	return user, nil
}
