package service

import (
	"github.com/gaganchawara/loans/internal/loan/interfaces"
)

type service struct {
	repo interfaces.Repository
}

func NewService(repo interfaces.Repository) interfaces.Service {
	return &service{
		repo: repo,
	}
}
