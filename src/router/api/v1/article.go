package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct {}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create(c *gin.Context) {
	c.JSON(200, gin.H{ "message": "success" })
}
