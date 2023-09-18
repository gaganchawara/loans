package server

import (
	"github.com/gaganchawara/loans/internal/loan/interfaces"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
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
