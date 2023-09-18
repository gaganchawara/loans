package server

import (
	"context"

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

func (s *server) GetLoans(ctx context.Context, request *loansv1.GetLoansRequest) (*loansv1.GetLoansResponse, error) {
	span, ctx := ot.StartSpanFromContext(ctx, "loan.server.GetLoans")
	defer span.Finish()

	if loan, ierr := s.service.GetLoanById(ctx, request.LoanId); ierr != nil {
		return nil, ierr
	} else {
		return &loansv1.GetLoansResponse{Loan: loan.Proto()}, nil
	}
}
