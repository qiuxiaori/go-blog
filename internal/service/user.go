package service

import (
	"github.com/qiuxiaori/go-blog/internal/dao"
)

type CreateUserReq struct {
	Name string `form:"name" binding:"required,min=2"`
}

func CreateUser(req *CreateUserReq) any {
	println("this is ", req.Name)
	return dao.CreateUser()
}
