package domain

import (
	"github.com/lib/pq"
)

type Expense struct {
	ID     uint `gorm:"primaryKey" param:"id"`
	Title  string
	Amount float64
	Note   string
	Tags   pq.StringArray `gorm:"type:text[]"`
}
