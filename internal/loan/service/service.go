package service

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/entity"
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

func (s *service) GetLoanById(ctx context.Context, loanId string) (*entity.Loan, errors.Error) {
	return s.repo.LoadLoan(ctx, loanId)
}
