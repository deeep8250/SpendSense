package services

import (
	"github.com/deeep8250/SpendSense/internal/repositories"
	"github.com/deeep8250/SpendSense/models"
)

type BudgetService struct {
	BudgetRepo *repositories.BudgetRepository
}

func NewBudgetService(Repo repositories.BudgetRepository) *BudgetService {
	return &BudgetService{
		BudgetRepo: &Repo,
	}
}

func (s *BudgetService) CreateBudgetService(budget models.Budget) error {

	err := s.BudgetRepo.CreateBudget(budget)
	if err != nil {
		return err
	}
	return nil
}

func (s *BudgetService) GetBudgetService(userID int) ([]models.Budget, error) {
	budgets, err := s.BudgetRepo.GetBudgets(userID)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}
func (s *BudgetService) SummaryRepoService(budget models.Budget) ([]models.InsightSummary, error) {
	summary, err := s.BudgetRepo.SummaryRepo(budget)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
func (s *BudgetService) TopMerchantService(userID int) ([]models.TopMerchant, error) {
	TopMerchant, err := s.BudgetRepo.TopMerchant(userID)
	if err != nil {
		return nil, err
	}
	return TopMerchant, nil
}
func (s *BudgetService) TrendService(userID int) ([]models.SpendingTrend, error) {
	Trending, err := s.BudgetRepo.Trend(userID)
	if err != nil {
		return nil, err
	}
	return Trending, nil
}
