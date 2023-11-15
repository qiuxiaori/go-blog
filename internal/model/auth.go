package model

import (
	"github.com/jinzhu/gorm"
	"github.com/qiuxiaori/go-blog/global"
	"github.com/qiuxiaori/go-blog/pkg/db"
)

type Auth struct {
	*db.BaseModel
	AppKey    string
	AppSecret string
}

func (a Auth) Get() (Auth, error) {
	var auth Auth

	db := global.DB.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
