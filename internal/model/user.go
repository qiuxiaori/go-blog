package model

import (
	"github.com/qiuxiaori/go-blog/global"
)

type User struct {
	// *db.BaseModel
	Name string `json:"name"`
}

func (u User) Create() error {
	println("this", u.Name)
	return global.DB.Create(&u).Error
}

func (u User) Get() (User, error) {
	var user User
	err := global.DB.Where("name = ? ", u.Name).First(&user).Error
	return user, err
}
