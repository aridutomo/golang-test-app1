package domain

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleUsecase interface {
	GetAll() ([]Article, error)
	GetByID(id int) (Article, error)
}
