package validation

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gaganchawara/loans/internal/enums/loanstatus"
	"github.com/gaganchawara/loans/internal/loan/aggregate"
	"github.com/gaganchawara/loans/internal/loan/entity"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

func ValidateGetLoanRequest(
	ctx context.Context, loanId string,
) errors.Error {
	if loanId == "" {
		return errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")).Report()
	}

	return nil
}

func ValidateApplyLoanRequest(
	ctx context.Context, request *loansv1.ApplyLoanRequest,
) errors.Error {
	if request.Amount < 100 {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("loan amount should be 1 INR at least")).Report()
	}
	if request.Term <= 0 {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("term should be greater than zero")).Report()
	}

	return nil
}

func ValidateApproveLoanRequest(
	ctx context.Context, request *loansv1.ApproveLoanRequest,
) errors.Error {
	if request.LoanId == "" {
		return errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")).Report()
	}

	return nil
}

func ValidateRejectLoanRequest(
	ctx context.Context, request *loansv1.RejectLoanRequest,
) errors.Error {
	if request.LoanId == "" {
		return errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")).Report()
	}

	return nil
}

func ValidateRepayLoanRequest(
	ctx context.Context, request *loansv1.RepayLoanRequest,
) errors.Error {
	if request.LoanId == "" {
		return errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")).Report()
	}
	if request.Amount < 100 {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("payment amount should be 1 INR at least")).Report()
	}

	return nil
}

func ValidateApproveLoanAgg(ctx context.Context, e *entity.Loan) errors.Error {
	if e.Status != loanstatus.Pending {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("approval not allowed for this loan")).Report()
	}

	return nil
}

func ValidateRejectLoanAgg(ctx context.Context, e *entity.Loan) errors.Error {
	if e.Status != loanstatus.Pending {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("rejection not allowed for this loan")).Report()
	}

	return nil
}

func ValidateRepayLoanAgg(ctx context.Context, agg *aggregate.LoanAgg, req *loansv1.RepayLoanRequest) errors.Error {
	if agg.Loan.Status == loanstatus.Paid {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("loan is already paid")).Report()
	} else if govalidator.IsIn(agg.Loan.Status.String(), loanstatus.Rejected.String(), loanstatus.Pending.String()) {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("payment not allowed for loan in this state")).Report()
	}

	if agg.GetTotalDueAmount() < req.Amount {
		return errors.New(ctx, errorcode.ValidationError,
			fmt.Errorf("paid mount greater than total due loan amount")).Report()
	}

	return nil
}
