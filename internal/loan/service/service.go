package service

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/aggregate"
	"github.com/gaganchawara/loans/internal/loan/factory"
	"github.com/gaganchawara/loans/internal/loan/interfaces"
	"github.com/gaganchawara/loans/pkg/errors"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

type service struct {
	repo interfaces.Repository
}

func NewService(repo interfaces.Repository) interfaces.Service {
	return &service{
		repo: repo,
	}
}

func (s service) ApplyLoan(ctx context.Context, req *loansv1.ApplyLoanRequest) (*aggregate.LoanAgg, errors.Error) {
	loanAgg, ierr := factory.Build(ctx, req)
	if ierr != nil {
		return nil, ierr
	}

	ierr = s.repo.SaveLoanAgg(ctx, loanAgg)
	if ierr != nil {
		return nil, ierr
	}

	return loanAgg, nil
}

func (s service) GetLoanAggById(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error) {
	return s.repo.LoadLoanAgg(ctx, loanId)
}
