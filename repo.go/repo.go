package repo

import (
	"gorm.io/gorm"
)

type Repo interface {
}

type ExpenseRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *ExpenseRepo {
	return &ExpenseRepo{db: db}
}
