package repositories

import (
	"fmt"
	"strings"

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

		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("CATEGORY already exists")
		}

		return err
	}
	return nil

}

func (r *CategoryRepository) GetCategoryByName(userID int, category_name string) (int, error) {

	var cat_id int
	query := `select id from categories where  name=$1 and (user_id=$2 or user_id is null)`
	err := r.db.Get(&cat_id, query, category_name, userID)
	if err != nil {
		return 0, err
	}

	return cat_id, nil

}
