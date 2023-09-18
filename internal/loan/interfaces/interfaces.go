package interfaces

import (
	"context"
	"github.com/gaganchawara/loans/internal/loan/aggregate"

	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
)

type Service interface {
	GetLoanAggById(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error)
}

type Repository interface {
	CreateLoan(ctx context.Context, loan *entity.Loan) errors.Error
	UpdateLoan(ctx context.Context, loan *entity.Loan) errors.Error
	LoadLoanAgg(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error)
	LoadLoan(ctx context.Context, loanId string) (*entity.Loan, errors.Error)
	LoadRepaymentsByLoanID(ctx context.Context, loanId string) ([]*entity.Repayment, errors.Error)
	LoadRepayment(ctx context.Context, id string) (*entity.Repayment, errors.Error)
}
