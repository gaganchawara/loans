package factory

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gaganchawara/loans/internal/helper"

	ctxkeys "github.com/gaganchawara/loans/internal/constants/ctx_keys"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/internal/loan/aggregate"
	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

func BuildLoan(ctx context.Context, request *loansv1.ApplyLoanRequest) (*aggregate.LoanAgg, errors.Error) {
	loan, ierr := entity.NewLoanEntity(ctx)
	if ierr != nil {
		return nil, ierr
	}

	userId, ok := ctx.Value(ctxkeys.UserID).(string)
	if !ok {
		return nil, errors.New(ctx, errorcode.InternalServerError,
			fmt.Errorf("user data not available in the context")).
			Report()
	}

	loan.UserId = userId
	loan.Amount = request.Amount
	loan.Term = request.Term

	return &aggregate.LoanAgg{Loan: loan}, nil
}

func BuildRepayments(_ context.Context, loan *entity.Loan) (*aggregate.LoanAgg, errors.Error) {
	var repayments []*entity.Repayment
	repaymentAmounts := helper.BreakAmount(loan.Amount, int(loan.Term))

	now := time.Now()

	for _, repaymentAmount := range repaymentAmounts {
		now = now.AddDate(0, 0, 7)
		repayment := entity.NewRepaymentEntity()
		repayment.LoanId = loan.Id
		repayment.Amount = repaymentAmount
		repayment.DueDate = sql.NullTime{Time: now, Valid: false}

		repayments = append(repayments, repayment)
	}

	return &aggregate.LoanAgg{Loan: loan, Repayments: repayments}, nil
}
