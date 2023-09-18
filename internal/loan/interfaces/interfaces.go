package interfaces

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
)

type Service interface {
	GetLoanById(ctx context.Context, loanId string) (*entity.Loan, errors.Error)
}

type Repository interface {
	CreateLoan(ctx context.Context, loan *entity.Loan) errors.Error
	UpdateLoan(ctx context.Context, loan *entity.Loan) errors.Error
	LoadLoan(ctx context.Context, loanId string) (*entity.Loan, errors.Error)
	LoadRepayment(ctx context.Context, id string) (*entity.Repayment, errors.Error)
}
