package interfaces

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/aggregate"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"

	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
)

type Service interface {
	ApplyLoan(ctx context.Context, req *loansv1.ApplyLoanRequest) (*aggregate.LoanAgg, errors.Error)
	ApproveLoan(ctx context.Context, req *loansv1.ApproveLoanRequest) (*aggregate.LoanAgg, errors.Error)
	RejectLoan(ctx context.Context, req *loansv1.RejectLoanRequest) (*aggregate.LoanAgg, errors.Error)
	RepayLoan(ctx context.Context, req *loansv1.RepayLoanRequest) (*aggregate.LoanAgg, errors.Error)
	GetLoanAggById(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error)
}

type Repository interface {
	SaveLoanAgg(ctx context.Context, loanAgg *aggregate.LoanAgg) errors.Error
	SaveLoan(ctx context.Context, loan *entity.Loan) errors.Error
	SaveRepayment(ctx context.Context, repayment *entity.Repayment) errors.Error
	LoadLoanAgg(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error)
	LoadLoan(ctx context.Context, loanId string) (*entity.Loan, errors.Error)
	LoadRepaymentsByLoanID(ctx context.Context, loanId string) ([]*entity.Repayment, errors.Error)
	LoadRepayment(ctx context.Context, id string) (*entity.Repayment, errors.Error)
}
