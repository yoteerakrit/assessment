package usecase

import (
	"github.com/yoteerakrit/assessment/repo.go"
)

type UseCase interface {
}

type useCase struct {
	repo repo.Repo
}
