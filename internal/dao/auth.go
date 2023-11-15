package dao

import "github.com/qiuxiaori/go-blog/internal/model"

type Auth struct{}

// func NewAuth () {
// 	A
// }
func (a Auth) Get(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get()
}
