package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuxiaori/go-blog/pkg/app"
)

type Article struct {
	Title string
}

func NewArticle() Article {
	return Article{}
}

// @Summary 创建文章
// @Produce  json
// @Param title body string true "文章标题" minlength(1)
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {

	// param := service.CreateArticleRequest{}
	// // response := app.NewResponse(c)
	// valid, errs := app.BindAndValid(c, &param)
	// if !valid {
	// 	println("this is errs", errs)
	// 	// global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
	// 	// response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	// 	return
	// }
	app.NewResponse(c).ToResponse(Article{"测试"})
}
