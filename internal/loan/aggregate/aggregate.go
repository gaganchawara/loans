package aggregate

import (
	"context"
	"sort"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"

	"github.com/gaganchawara/loans/internal/enums/repaymentstatus"
	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

type LoanAgg struct {
	Loan       *entity.Loan
	Repayments []*entity.Repayment
}

func (agg *LoanAgg) RepayAmount(ctx context.Context, paidAmount int64) errors.Error {
	SortRepaymentsByDueDate(agg.Repayments)

	leftAmount := paidAmount

	for _, repayment := range agg.Repayments {
		if leftAmount == 0 {
			break
		}
		if repayment.Status == repaymentstatus.Paid {
			continue
		}

		if repayment.Amount-repayment.PaidAmount > leftAmount {
			repayment.Status = repaymentstatus.PartiallyPaid
			repayment.PaidAmount += leftAmount
			leftAmount = 0
		} else {
			leftAmount -= repayment.Amount - repayment.PaidAmount
			repayment.Status = repaymentstatus.Paid
			repayment.PaidAmount = repayment.Amount
		}
	}

	if agg.GetTotalDueAmount() == 0 {
		agg.Loan.Status = loanstatus.Paid
	} else {
		agg.Loan.Status = loanstatus.PartiallyPaid
	}

	return nil
}

func (agg *LoanAgg) GetTotalDueAmount() int64 {
	var dueAmount int64
	for _, repayment := range agg.Repayments {
		dueAmount += repayment.Amount - repayment.PaidAmount
	}

	return dueAmount
}

// SortRepaymentsByDueDate sorts a slice of Repayment entities by their DueDate in ascending order.
func SortRepaymentsByDueDate(repayments []*entity.Repayment) {
	sort.Slice(repayments, func(i, j int) bool {
		return repayments[i].DueDate.Time.Before(repayments[j].DueDate.Time)
	})
}

func (agg *LoanAgg) Proto() *loansv1.LoansResponse {
	var repaymentProtos []*loansv1.Repayment
	for _, repayment := range agg.Repayments {
		repaymentProtos = append(repaymentProtos, repayment.Proto())
	}

	return &loansv1.LoansResponse{
		Loan:       agg.Loan.Proto(),
		Repayments: repaymentProtos,
	}
}
