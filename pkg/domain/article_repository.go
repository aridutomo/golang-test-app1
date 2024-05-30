package domain

type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetByID(id int) (Article, error)
}
