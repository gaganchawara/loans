package server

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/tracecode"
	"github.com/gaganchawara/loans/pkg/logger"

	"github.com/gaganchawara/loans/internal/loan/interfaces"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
	ot "github.com/opentracing/opentracing-go"
)

type server struct {
	loansv1.UnimplementedLoansAPIServer
	service interfaces.Service
}

func NewServer(service interfaces.Service) *server {
	return &server{
		service: service,
	}
}

func (s *server) ApplyLoan(ctx context.Context, request *loansv1.ApplyLoanRequest) (*loansv1.LoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.ApplyLoan")
	defer span.Finish()

	logger.Get(ctx).WithFields(map[string]interface{}{
		"amount": request.Amount,
		"term":   request.Term,
	}).Info(tracecode.ApplyLoanRequest)

	if agg, ierr := s.service.ApplyLoan(ctx, request); ierr != nil {
		return nil, ierr
	} else {
		return agg.Proto(), nil
	}
}

func (s *server) ApproveLoan(ctx context.Context, request *loansv1.ApproveLoanRequest) (*loansv1.LoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.ApproveLoan")
	defer span.Finish()

	logger.Get(ctx).WithFields(map[string]interface{}{
		"loan_id": request.LoanId,
	}).Info(tracecode.ApproveLoanRequest)

	if agg, ierr := s.service.ApproveLoan(ctx, request); ierr != nil {
		return nil, ierr
	} else {
		return agg.Proto(), nil
	}
}

func (s *server) RejectLoan(ctx context.Context, request *loansv1.RejectLoanRequest) (*loansv1.LoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.RejectLoan")
	defer span.Finish()

	logger.Get(ctx).WithFields(map[string]interface{}{
		"loan_id": request.LoanId,
	}).Info(tracecode.RejectLoanRequest)

	if agg, ierr := s.service.RejectLoan(ctx, request); ierr != nil {
		return nil, ierr
	} else {
		return agg.Proto(), nil
	}
}

func (s *server) RepayLoan(ctx context.Context, request *loansv1.RepayLoanRequest) (*loansv1.LoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.RepayLoan")
	defer span.Finish()

	logger.Get(ctx).WithFields(map[string]interface{}{
		"loan_id": request.LoanId,
		"amount":  request.Amount,
	}).Info(tracecode.RepayLoanRequest)

	if agg, ierr := s.service.RepayLoan(ctx, request); ierr != nil {
		return nil, ierr
	} else {
		return agg.Proto(), nil
	}
}

func (s *server) GetLoans(ctx context.Context, request *loansv1.GetLoanRequest) (*loansv1.LoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.GetLoans")
	defer span.Finish()

	logger.Get(ctx).WithFields(map[string]interface{}{
		"loan_id": request.LoanId,
	}).Info(tracecode.GetLoanRequest)

	if agg, ierr := s.service.GetLoanAggById(ctx, request.LoanId); ierr != nil {
		return nil, ierr
	} else {
		return agg.Proto(), nil
	}
}
