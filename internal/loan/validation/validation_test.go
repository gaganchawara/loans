package validation

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/internal/helper"
	"github.com/gaganchawara/loans/internal/loan/aggregate"
	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
)

func TestValidateGetLoanRequest(t *testing.T) {
	ctx := context.Background()

	type args struct {
		loanId string
	}
	tests := []struct {
		name string
		args args
		want errors.Error
	}{
		{
			name: "success",
			args: args{
				loanId: "LoanId",
			},
			want: nil,
		},
		{
			name: "failure",
			args: args{
				loanId: "",
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateGetLoanRequest(ctx, tt.args.loanId)
			helper.AssertEqualError(t, tt.want, got)
		})
	}
}

func TestValidateApplyLoanRequest(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request *loansv1.ApplyLoanRequest
	}
	tests := []struct {
		name string
		args args
		want errors.Error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: &loansv1.ApplyLoanRequest{
					Amount: 10000,
					Term:   10,
				},
			},
			want: nil,
		},
		{
			name: "fail_amount",
			args: args{
				ctx: ctx,
				request: &loansv1.ApplyLoanRequest{
					Amount: 99,
					Term:   10,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan amount should be 1 INR at least")),
		},
		{
			name: "fail_term",
			args: args{
				ctx: ctx,
				request: &loansv1.ApplyLoanRequest{
					Amount: 10000,
					Term:   0,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("term should be greater than zero")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateApplyLoanRequest(tt.args.ctx, tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateApplyLoanRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateRepayLoanRequest(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request *loansv1.RepayLoanRequest
	}
	tests := []struct {
		name string
		args args
		want errors.Error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: &loansv1.RepayLoanRequest{
					LoanId: "randomLoanId",
					Amount: 10000,
				},
			},
			want: nil,
		},
		{
			name: "fail_amount",
			args: args{
				ctx: ctx,
				request: &loansv1.RepayLoanRequest{
					LoanId: "randomLoanId",
					Amount: 0,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("payment amount should be 1 INR at least")),
		},
		{
			name: "fail_loanId",
			args: args{
				ctx: ctx,
				request: &loansv1.RepayLoanRequest{
					Amount: 10000,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan id required")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateRepayLoanRequest(tt.args.ctx, tt.args.request)
			helper.AssertEqualError(t, tt.want, got)
		})
	}
}

func TestValidateRepayLoanAgg(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		agg *aggregate.LoanAgg
		req *loansv1.RepayLoanRequest
	}
	tests := []struct {
		name string
		args args
		want errors.Error
	}{
		{
			name: "loan_paid_failure",
			args: args{
				ctx: ctx,
				agg: &aggregate.LoanAgg{
					Loan: &entity.Loan{
						Id:     "randomLoanId",
						Status: loanstatus.Paid,
					},
				},
				req: &loansv1.RepayLoanRequest{
					LoanId: "randomLoanId",
					Amount: 10000,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("loan is already paid")),
		},
		{
			name: "loan_rejected_failure",
			args: args{
				ctx: ctx,
				agg: &aggregate.LoanAgg{
					Loan: &entity.Loan{
						Id:     "randomLoanId",
						Status: loanstatus.Rejected,
					},
				},
				req: &loansv1.RepayLoanRequest{
					LoanId: "randomLoanId",
					Amount: 10000,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("payment not allowed for loan in this state")),
		},
		{
			name: "loan_success",
			args: args{
				ctx: ctx,
				agg: &aggregate.LoanAgg{
					Loan: &entity.Loan{
						Id:     "randomLoanId",
						Amount: 100000,
						Status: loanstatus.Rejected,
					},
					Repayments: []*entity.Repayment{
						{
							Amount: 100000,
						},
					},
				},
				req: &loansv1.RepayLoanRequest{
					LoanId: "randomLoanId",
					Amount: 100000,
				},
			},
			want: errors.New(ctx, errorcode.ValidationError, fmt.Errorf("payment not allowed for loan in this state")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateRepayLoanAgg(tt.args.ctx, tt.args.agg, tt.args.req)
			helper.AssertEqualError(t, tt.want, got)
		})
	}
}
