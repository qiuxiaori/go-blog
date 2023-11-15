package service

import (
	"github.com/qiuxiaori/go-blog/internal/dao"
	"github.com/qiuxiaori/go-blog/internal/model"
)

type Auth struct{}

func (a Auth) Get(appKey, appSecret string) (model.Auth, error) {
	return dao.Auth{}.Get(appKey, appSecret)
}
