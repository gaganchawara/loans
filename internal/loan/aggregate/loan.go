package aggregate

import (
	"github.com/gaganchawara/loans/internal/loan/entity"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

type LoanAgg struct {
	Loan       *entity.Loan
	Repayments []*entity.Repayment
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
