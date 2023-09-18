package service

import (
	"context"
	"github.com/gaganchawara/loans/internal/loan/aggregate"

	"github.com/gaganchawara/loans/internal/loan/interfaces"
	"github.com/gaganchawara/loans/pkg/errors"
)

type service struct {
	repo interfaces.Repository
}

func NewService(repo interfaces.Repository) interfaces.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetLoanAggById(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error) {
	return s.repo.LoadLoanAgg(ctx, loanId)
}
