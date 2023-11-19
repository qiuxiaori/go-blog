package dao

import "github.com/qiuxiaori/go-blog/internal/model"

func GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get()
}
