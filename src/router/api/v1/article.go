package v1

import (
	// "github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

// todo 怎么获取私有参数
type Article struct {
	Title string `form:"title" binding:"required"`
}

func NewArticle() Article {
	return Article{}
}

// @Summary 创建文章
// @Produce  json
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	app.NewResponse(c).ToResponse(Article{"测试"})
}
