package repository

import "testapp1/pkg/domain"

type articleRepository struct {
	articles []domain.Article
}

func (r *articleRepository) GetAll() ([]domain.Article, error) {
	return r.articles, nil
}

func NewArticleRepository() domain.ArticleRepository {
	return &articleRepository{
		articles: []domain.Article{
			{ID: 1, Title: "Title 1", Content: "Body 1"},
			{ID: 2, Title: "Title 2", Content: "Body 2"},
		},
	}
}

func (r *articleRepository) GetByID(id int) (domain.Article, error) {
	for _, article := range r.articles {
		if article.ID == id {
			return article, nil
		}
	}
	return domain.Article{}, nil
}
