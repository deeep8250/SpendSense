package repositories

import (
	"github.com/deeep8250/SpendSense/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(Db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: Db,
	}
}

func (r *AuthRepository) RegisterUser(user models.User) error {
	query := `insert into users(name,email,hashed_password) values($1,$2,$3)`
	_, err := r.db.Exec(query, user.Name, user.Email, user.HashedPassword)
	if err != nil {
		return err
	}
	return nil
}
