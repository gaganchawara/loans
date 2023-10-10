package service

import (
	"context"
	"time"

	"github.com/gaganchawara/loans/internal/loan/validation"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"

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
	if ierr := validation.ValidateApplyLoanRequest(ctx, req); ierr != nil {
		return nil, ierr
	}

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
	if ierr := validation.ValidateApproveLoanRequest(ctx, req); ierr != nil {
		return nil, ierr
	}

	loan, ierr := s.repo.LoadLoan(ctx, req.LoanId)
	if ierr != nil {
		return nil, ierr
	}

	if ierr := validation.ValidateApproveLoanAgg(ctx, loan); ierr != nil {
		return nil, ierr
	}

	adminId, ierr := iam.GetAdminId(ctx)
	if ierr != nil {
		return nil, ierr
	}

	loan.ApprovedBy = adminId
	now := time.Now()
	loan.DisbursedAt = &now
	loan.Status = loanstatus.Approved

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

func (s service) RejectLoan(ctx context.Context, req *loansv1.RejectLoanRequest) (*aggregate.LoanAgg, errors.Error) {
	if ierr := validation.ValidateRejectLoanRequest(ctx, req); ierr != nil {
		return nil, ierr
	}

	agg, ierr := s.repo.LoadLoanAgg(ctx, req.LoanId)
	if ierr != nil {
		return nil, ierr
	}

	if ierr := validation.ValidateRejectLoanAgg(ctx, agg.Loan); ierr != nil {
		return nil, ierr
	}

	agg.Loan.Status = loanstatus.Rejected

	ierr = s.repo.SaveLoanAgg(ctx, agg)
	if ierr != nil {
		return nil, ierr
	}

	return agg, nil
}

func (s service) RepayLoan(ctx context.Context, req *loansv1.RepayLoanRequest) (*aggregate.LoanAgg, errors.Error) {
	if ierr := validation.ValidateRepayLoanRequest(ctx, req); ierr != nil {
		return nil, ierr
	}

	agg, ierr := s.repo.LoadLoanAgg(ctx, req.LoanId)
	if ierr != nil {
		return nil, ierr
	}

	if ierr := validation.ValidateRepayLoanAgg(ctx, agg, req); ierr != nil {
		return nil, ierr
	}

	if ierr = iam.ValidateAccessToUser(ctx, agg.Loan.UserId); ierr != nil {
		return nil, ierr
	}

	if ierr = agg.RepayAmount(ctx, req.Amount); ierr != nil {
		return nil, ierr
	}

	ierr = s.repo.SaveLoanAgg(ctx, agg)
	if ierr != nil {
		return nil, ierr
	}

	return agg, nil
}

func (s service) GetLoanAggById(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error) {
	if ierr := validation.ValidateGetLoanRequest(ctx, loanId); ierr != nil {
		return nil, ierr
	}

	agg, ierr := s.repo.LoadLoanAgg(ctx, loanId)
	if ierr != nil {
		return nil, ierr
	}

	if ierr = iam.ValidateAccessToUser(ctx, agg.Loan.UserId); ierr != nil {
		return nil, ierr
	}

	return agg, nil
}
