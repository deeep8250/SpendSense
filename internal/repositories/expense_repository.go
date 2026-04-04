package repositories

import (
	"errors"
	"fmt"

	"github.com/deeep8250/SpendSense/models"
	"github.com/jmoiron/sqlx"
)

type ExpenseRepository struct {
	db *sqlx.DB
}

func NewExpenseRepository(Db *sqlx.DB) *ExpenseRepository {
	return &ExpenseRepository{

		db: Db,
	}
}

func (e *ExpenseRepository) CreateExpense(expense models.Expense) error {
	query := `insert into expenses(amount,merchant,category_id,description,user_id,source,date) values($1,$2,$3,$4,$5,$6,$7)`
	_, err := e.db.Exec(query, expense.Amount, expense.Merchant, expense.CategoryID, expense.Description, expense.UserID, expense.Source, expense.Date)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExpenseRepository) GetAllExpenses(userID int) ([]models.Expense, error) {
	var expenses []models.Expense
	query := `select * from expenses where user_id=$1`
	err := e.db.Select(&expenses, query, userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil

}

func (e *ExpenseRepository) GetSingleExpense(userID, expense_id int) (*models.Expense, error) {
	var expense models.Expense
	query := `select * from expenses where user_id =$1 and id=$2`
	err := e.db.Get(&expense, query, userID, expense_id)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (e *ExpenseRepository) RemoveExpense(user_id, expense_id int) (string, error) {
	query := `delete from expenses where user_id=$1 and id=$2`
	result, err := e.db.Exec(query, user_id, expense_id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("data isnt deleted")
	}

	return "deleted", nil
}

func (e *ExpenseRepository) GetAllExpensesByVariousFilters(userID int, categoryID, source, Date string) ([]models.Expense, error) {

	query := `select *from expenses where user_id=$1`
	args := []any{userID}
	argsCount := 1
	if categoryID != "" {
		argsCount++
		query += fmt.Sprintf(" and category_id=$%d", argsCount)

		args = append(args, categoryID)

	}

	if source != "" {
		argsCount++
		query += fmt.Sprintf(" and source=$%d", argsCount)

		args = append(args, source)

	}

	if Date != "" {
		argsCount++
		query += fmt.Sprintf(" and date=$%d", argsCount)

		args = append(args, Date)

	}

	var expenses []models.Expense
	err := e.db.Select(&expenses, query, args...)
	if err != nil {
		return nil, err
	}
	return expenses, nil

}
