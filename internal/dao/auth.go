package dao

import "github.com/qiuxiaori/go-blog/internal/model"

func GetUser(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get()
}
