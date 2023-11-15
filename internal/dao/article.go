package dao

type ArticleSwagger struct {
	List []*string
}

type Article struct {
	Title string
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create() Article {
	return Article{}
}
