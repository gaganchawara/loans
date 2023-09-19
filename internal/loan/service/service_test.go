package service

import (
	"context"
	"testing"

	ctxkeys "github.com/gaganchawara/loans/internal/constants/ctx_keys"
	"github.com/gaganchawara/loans/internal/enums/accounttype"
	"github.com/gaganchawara/loans/internal/enums/loanstatus"
	"github.com/gaganchawara/loans/internal/loan/interfaces"
	"github.com/gaganchawara/loans/internal/loan/mock"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"
	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/suite"
)

type serviceTestSuit struct {
	ctx      context.Context
	mockRepo *mock.MockRepository

	svc interfaces.Service

	suite.Suite
}

func TestLoansService(t *testing.T) {
	suite.Run(t, new(serviceTestSuit))
}

func (suite *serviceTestSuit) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockRepo = mock.NewMockRepository(ctrl)
	suite.svc = NewService(suite.mockRepo)
}

func (suite *serviceTestSuit) Test_ApplyLoan() {
	suite.ctx = context.WithValue(context.Background(), ctxkeys.AccountType, accounttype.User)
	suite.ctx = context.WithValue(suite.ctx, ctxkeys.UserID, "randomUserId")

	suite.mockRepo.EXPECT().SaveLoanAgg(gomock.Any(), gomock.Any()).Return(nil)

	loanAgg, ierr := suite.svc.ApplyLoan(suite.ctx, &loansv1.ApplyLoanRequest{
		Amount: 10000,
		Term:   2,
	})

	suite.Assert().Nil(ierr)
	suite.Assert().Nil(loanAgg.Repayments)
	suite.Assert().Equal(int64(10000), loanAgg.Loan.Amount)
	suite.Assert().Equal(int32(2), loanAgg.Loan.Term)
	suite.Assert().Equal(loanstatus.Pending, loanAgg.Loan.Status)
	suite.Assert().Equal("randomUserId", loanAgg.Loan.UserId)
}
