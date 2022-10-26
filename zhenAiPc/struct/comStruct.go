package _struct

import "gorm.io/gorm"

type UserXyy struct {
	gorm.Model
	Name  string
	Href  string
	Photo string

	House      string //结婚
	Age        string
	Xinzuo     string //星座
	Height     string
	Weight     string
	Occupation string //工作地
	Income     string //收入
	Work       string //职位
	Education  string //教育学历
	Gender     string //性别
	Marriage   string //结婚
	Car        string //车
	Habitation string //居住地
}

type CityList struct {
	CityList []CityListName
}

type CityListName struct {
	Name string
	Href string
}

func NewFuncParser(href string, name string) *CityListName {
	return &CityListName{
		Name: name,
		Href: href,
	}
}
