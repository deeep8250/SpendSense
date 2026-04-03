package services

import (
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/models"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(Repo *repositories.CategoryRepository) *CategoryService {

	return &CategoryService{repo: Repo}
}

func (s *CategoryService) GetCategories(userID int) ([]models.Category, error) {

	categories, err := s.repo.GetCategories(userID)
	if err != nil {
		return nil, err
	}
	return categories, nil

}

func (s *CategoryService) CreateCategory(category models.Category) error {

	err := s.repo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil

}
