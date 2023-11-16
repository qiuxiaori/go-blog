package service

import (
	"github.com/qiuxiaori/go-blog/internal/dao"
	"github.com/qiuxiaori/go-blog/internal/model"
)

func GetAuth(appKey, appSecret string) (model.Auth, error) {
	return dao.GetUser(appKey, appSecret)
}
