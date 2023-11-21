package dao

import "github.com/qiuxiaori/go-blog/internal/model"

func CreateUser(name string) error {
	user := model.User{Name: name}
	println("this is user 1", name, user.Name)
	return user.Create()
}

func GetUser(name string) (model.User, error) {
	user := model.User{Name: name}
	return user.Get()
}
