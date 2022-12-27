package repo

import (
	"gorm.io/gorm"
)

type Repo interface {
}

type ExpenseRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ExpenseRepo {
	return &ExpenseRepo{db: db}
}
