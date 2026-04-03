package services

import (
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/models"
)

type ExpenseService struct {
	repo *repositories.ExpenseRepository
}

func NewExpenseService(Repo *repositories.ExpenseRepository) *ExpenseService {

	return &ExpenseService{repo: Repo}
}

func (s *ExpenseService) GetExpenses(userID int) ([]models.Expense, error) {
	expenses, err := s.repo.GetAllExpenses(userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *ExpenseService) GetExpensesByID(userID, expenseID int) (*models.Expense, error) {
	expense, err := s.repo.GetSingleExpense(userID, expenseID)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (s *ExpenseService) DeleteExpense(userID, expenseID int) (string, error) {
	result, err := s.repo.RemoveExpense(userID, expenseID)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *ExpenseService) CreateExpense(expense models.Expense) error {
	err := s.repo.CreateExpense(expense)
	if err != nil {
		return err
	}
	return nil
}
