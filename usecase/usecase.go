package usecase

import (
	"github.com/yoteerakrit/assessment/repo"
)

type UseCase interface {
}

type useCase struct {
	repo repo.Repo
}
