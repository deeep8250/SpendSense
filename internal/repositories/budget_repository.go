package repositories

import (
	"fmt"
	"strings"

	"github.com/deeep8250/SpendSense/models"
	"github.com/jmoiron/sqlx"
)

type BudgetRepository struct {
	db *sqlx.DB
}

func NewBudgetRepository(DB *sqlx.DB) *BudgetRepository {

	return &BudgetRepository{db: DB}
}

func (r *BudgetRepository) CreateBudget(budget models.Budget) error {
	query := `insert into budgets(user_id,category_id,amount,month,year) values($1,$2,$3,$4,$5)`
	_, err := r.db.Exec(query, budget.UserID, budget.CategoryID, budget.Amount, budget.Month, budget.Year)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("budget already exists for this category and month")
		}

		return err
	}
	return nil
}

func (r *BudgetRepository) GetBudgets(userID int) ([]models.Budget, error) {
	var budgets []models.Budget
	query := `select * from budgets where user_id=$1`
	err := r.db.Select(&budgets, query, userID)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepository) BudgetAlert(userID int) ([]models.BudgetAlert, error) {

	var BudgetAlertHold []models.BudgetAlert

	query := `SELECT b.id, b.category_id, b.amount AS budget_amount, COALESCE(SUM(e.amount), 0)
	 AS total_spent FROM budgets b 
	 LEFT JOIN expenses e ON b.category_id = e.category_id 
	 AND b.user_id = e.user_id
	  AND EXTRACT(MONTH FROM e.date) = b.month 
	  AND EXTRACT(YEAR FROM e.date) = b.year 
	  WHERE b.user_id = $1 
	  GROUP BY b.id, b.category_id, b.amount 
	  HAVING COALESCE(SUM(e.amount), 0) > b.amount`

	err := r.db.Select(&BudgetAlertHold, query, userID)
	if err != nil {
		return nil, err
	}
	return BudgetAlertHold, err
}

func (r *BudgetRepository) SummaryRepo(userId, month, year int) ([]models.InsightSummary, error) {
	var Summary []models.InsightSummary
	query := `SELECT c.name AS category, COALESCE(SUM(e.amount), 0) AS total 
	          FROM expenses e JOIN categories c ON e.category_id = c.id 
			  WHERE e.user_id = $1 
			  AND EXTRACT(MONTH FROM e.date) = $2 
			  AND EXTRACT(YEAR FROM e.date) = $3 
			  GROUP BY c.name`

	err := r.db.Select(&Summary, query, userId, month, year)
	if err != nil {
		return nil, err
	}

	return Summary, nil

}

func (r *BudgetRepository) TopMerchant(userID int) ([]models.TopMerchant, error) {
	var TopMerchantList []models.TopMerchant
	query := `SELECT merchant, SUM(amount) AS total FROM expenses WHERE user_id = $1 GROUP BY merchant ORDER BY total DESC LIMIT 5`
	err := r.db.Select(&TopMerchantList, query, userID)
	if err != nil {
		return nil, err
	}
	return TopMerchantList, nil
}

func (r *BudgetRepository) Trend(userID int) ([]models.SpendingTrend, error) {

	var TrendList []models.SpendingTrend
	query := `SELECT EXTRACT(MONTH FROM date) AS month, 
	         EXTRACT(YEAR FROM date) AS year, SUM(amount) AS total 
			  FROM expenses WHERE user_id = $1 
			  GROUP BY month, year 
			  ORDER BY year, month`

	err := r.db.Select(&TrendList, query, userID)
	if err != nil {
		return nil, err
	}
	return TrendList, nil

}
