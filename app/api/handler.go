package api

import "github.com/yoteerakrit/assessment/usecase"

type expenseHandler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *expenseHandler {
	return &expenseHandler{uc: uc}
}
