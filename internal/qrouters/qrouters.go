package qrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter"
	"github.com/qiuxiaori/go-blog/pkg/qiurouter/user"
)

func Init(gin *gin.Engine) {
	qiurouter.New(gin).Register(user.Create())
}
