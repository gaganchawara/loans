package service

import (
	"context"

	"github.com/gaganchawara/loans/internal/loans/interfaces"
)

type service struct {
	repo interfaces.Repository
}

func NewService(ctx context.Context, repo interfaces.Repository) interfaces.Service {
	return &service{
		repo: repo,
	}
}
