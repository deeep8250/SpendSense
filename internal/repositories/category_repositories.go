package repositories

import (
	"github.com/deeep8250/SpendSense/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(Db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{

		db: Db,
	}
}

func (r *CategoryRepository) GetCategories(userID int) ([]models.Category, error) {
	var categories []models.Category
	query := `select * from categories where user_id=$1 OR user_id is null`
	err := r.db.Select(&categories, query, userID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CreateCategory(category models.Category) error {
	query := `insert into categories (name,user_id) values($1,$2)`
	_, err := r.db.Exec(query, category.Name, category.UserID)
	if err != nil {
		return err
	}
	return nil

}
