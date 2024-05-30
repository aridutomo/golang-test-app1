package usecase

import "testapp1/pkg/domain"

type articleUsecase struct {
	repo domain.ArticleRepository
}

func (u *articleUsecase) GetAll() ([]domain.Article, error) {
	return u.repo.GetAll()
}

func NewArticleUsecase(repo domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{
		repo: repo,
	}
}

func (u *articleUsecase) GetByID(id int) (domain.Article, error) {
	return u.repo.GetByID(id)
}
