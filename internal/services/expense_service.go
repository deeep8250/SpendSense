package services

import (
	"encoding/json"

	parses "github.com/deeep8250/SpendSense/internal/parser"
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/models"
)

type ExpenseService struct {
	repo         *repositories.ExpenseRepository
	CategoryRepo *repositories.CategoryRepository
}

func NewExpenseService(Repo *repositories.ExpenseRepository, CategoryRepo *repositories.CategoryRepository) *ExpenseService {
	return &ExpenseService{
		repo:         Repo,
		CategoryRepo: CategoryRepo,
	}
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

func (s *ExpenseService) GetAllExpensesByDifferentFilter(userID int, categoryID, source, Date string) ([]models.Expense, error) {
	expenses, err := s.repo.GetAllExpensesByVariousFilters(userID, categoryID, source, Date)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *ExpenseService) ParseAiExpense(response string, userID int) error {

	//get the all available categories from the user id
	categories, err := s.CategoryRepo.GetCategories(userID)
	if err != nil {
		return err
	}

	// store them into a string slice through loop
	var categoryesForAiParser []string
	for _, category := range categories {
		categoryesForAiParser = append(categoryesForAiParser, category.Name)
	}

	// pass it to the ai parser it will return json
	AiParserResponse, err := parses.AiParser(response, categoryesForAiParser)
	if err != nil {
		return err
	}
	//unmarshal the ai response
	var FinalValue models.AiParserResponseHolder
	err = json.Unmarshal([]byte(AiParserResponse), &FinalValue)
	if err != nil {
		return err
	}

	// checking that ai returned values are valid or not
	err = parses.ValidateParsedExpense(FinalValue, categoryesForAiParser)
	if err != nil {
		return err
	}
	//get the category id of returned category name from ai parser
	category_id, err := s.CategoryRepo.GetCategoryByName(userID, FinalValue.Category)
	if err != nil {
		return err
	}

	// pass it to the repo
	FinalValue.Source = "AI"
	err = s.repo.ExpenseParser(FinalValue, category_id, userID)
	if err != nil {
		return err
	}
	return nil

}
