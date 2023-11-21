package service

import (
	"github.com/qiuxiaori/go-blog/internal/dao"
	"github.com/qiuxiaori/go-blog/internal/model"
)

type UserCreateReq struct {
	Name string `form:"name" binding:"required,min=1"`
}

type UserDetailReq struct {
	Name string `form:"name" binding:"required`
}

func CreateUser(req *UserCreateReq) error {
	return dao.CreateUser(req.Name)
}

func GetUser(req *UserDetailReq) (model.User, error) {
	return dao.GetUser(req.Name)
}
