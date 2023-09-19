package service

import (
	"context"
	"time"

	"github.com/gaganchawara/loans/internal/iam"

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
	loanAgg, ierr := factory.BuildLoan(ctx, req)
	if ierr != nil {
		return nil, ierr
	}

	ierr = s.repo.SaveLoanAgg(ctx, loanAgg)
	if ierr != nil {
		return nil, ierr
	}

	return loanAgg, nil
}

func (s service) ApproveLoan(ctx context.Context, req *loansv1.ApproveLoanRequest) (*aggregate.LoanAgg, errors.Error) {
	loan, ierr := s.repo.LoadLoan(ctx, req.LoanId)
	if ierr != nil {
		return nil, ierr
	}

	adminId, ierr := iam.GetAdminId(ctx)
	if ierr != nil {
		return nil, ierr
	}

	loan.ApprovedBy = adminId
	now := time.Now()
	loan.DisbursedAt = &now

	loanAgg, ierr := factory.BuildRepayments(ctx, loan)
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
