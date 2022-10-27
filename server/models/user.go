package models

import (
	gorm2 "gorm.io/gorm"
)

type UserXyy struct {
	gorm2.Model
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Href  string `json:"href"`

	Gender     string `json:"gender"`
	Habitation string `json:"habitation"`
	Age        string `json:"age"`
	Education  string `json:"education"`
	House      string `json:"house"`
	Height     string `json:"height"`
}

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Height string `json:"height"`
}
