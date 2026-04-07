package models

import "time"

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type User struct {
	Id             int       `db:"id"`
	Name           string    `db:"name"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAT      time.Time `db:"created_at"`
}

type Category struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	UserID    *int      `db:"user_id" json:"user_id" `
	CreatedAT time.Time `db:"created_at" json:"created_at"`
}

type Expense struct {
	Id          int       `db:"id" json:"id"`
	Amount      float64   `db:"amount" json:"amount" binding:"required,min=1.00"`
	Merchant    string    `db:"merchant" json:"merchant" binding:"required"`
	CategoryID  int       `db:"category_id" json:"category_id" binding:"required"`
	Description string    `db:"description" json:"description" binding:"required"`
	UserID      int       `db:"user_id" json:"user_id"`
	Date        string    `db:"date" json:"date" binding:"required"`
	Source      string    `db:"source" json:"source"`
	CreatedAT   time.Time `db:"created_at" json:"created_at"`
}

type Budget struct {
	Id         int       `db:"id" json:"id"`
	UserID     int       `db:"user_id" json:"user_id"`
	CategoryID int       `db:"category_id" json:"category_id" binding:"required"`
	Amount     float64   `db:"amount" json:"amount" binding:"required"`
	Month      int       `db:"month" json:"month" binding:"required"`
	Year       int       `db:"year" json:"year" binding:"required"`
	CreatedAT  time.Time `db:"created_at" json:"created_at"`
}

type AiParserResponseHolder struct {
	Amount      int
	Merchant    string
	Category    string
	Description string
	Source      string
	Date        string
}

type BudgetAlert struct {
	ID           int     `db:"id" json:"id"`
	CategoryID   int     `db:"category_id" json:"category_id"`
	BudgetAmount float64 `db:"budget_amount" json:"budget_amount"`
	TotalSpent   float64 `db:"total_spent" json:"total_spent"`
}

type InsightSummary struct {
	Category string  `db:"category" json:"category"`
	Total    float64 `db:"total" json:"total"`
}

type TopMerchant struct {
	Merchant string  `db:"merchant" json:"merchant"`
	Total    float64 `db:"total" json:"total"`
}

type SpendingTrend struct {
	Month int     `db:"month" json:"month"`
	Year  int     `db:"year" json:"year"`
	Total float64 `db:"total" json:"total"`
}
