package service

type CreateArticleRequest struct {
	Title string `form:"title" binding:"require,min=1,max=100"`
}

type Article struct {
	Title string
}

func NewArticle() Article {
	return Article{}
}

func CreateArticle() Article {
	return Article{}
}

// func (svc *Service) CreateArticle {
// 	return svc.model
// }
