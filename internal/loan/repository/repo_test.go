package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gaganchawara/loans/internal/helpertest"
	"github.com/gaganchawara/loans/internal/loan/interfaces"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbKeyType string

const (
	DbCtxKey DbKeyType = "Db"
)

type repositoryTestSuit struct {
	ctx     context.Context
	gormDB  *gorm.DB
	sqlmock sqlmock.Sqlmock
	repo    interfaces.Repository

	loanColumns      []string
	repaymentColumns []string
	suite.Suite
}

func TestLoansRepository(t *testing.T) {
	suite.Run(t, new(repositoryTestSuit))
}

func (suite *repositoryTestSuit) SetupTest() {
	suite.gormDB, suite.sqlmock = createMockDB(suite.T())
	suite.ctx = context.WithValue(context.TODO(), DbCtxKey, suite.gormDB)
	suite.repo = NewRepository(suite.gormDB)

	// defining the columns for mock DB
	suite.loanColumns = []string{
		"id", "user_id", "amount", "term", "status", "approved_by", "disbursed_at", "created_at", "updated_at",
		"deleted_at",
	}
	suite.repaymentColumns = []string{
		"id", "loan_id", "amount", "paid_amount", "status", "due_date", "created_at", "updated_at", "deleted_at",
	}
}

func createMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	conn, mockdb, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Create a new GORM DB connection using the mock
	gormDB, err := gorm.Open(getGormDialectorForMock(conn))
	if err != nil {
		t.Fatalf("Failed to open GORM DB: %v", err)
	}

	return gormDB, mockdb
}

func getGormDialectorForMock(conn gorm.ConnPool) gorm.Dialector {
	return mysql.New(mysql.Config{Conn: conn, SkipInitializeWithVersion: true})
}

func (suite *repositoryTestSuit) Test_LoadRepayment() {
	repayment := helpertest.GetTestRepaymentEntity()

	repaymentRow := suite.sqlmock.NewRows(suite.repaymentColumns).
		AddRow(
			repayment.Id, repayment.LoanId, repayment.Amount, repayment.PaidAmount, []byte(repayment.Status.String()),
			repayment.DueDate, repayment.CreatedAt, repayment.UpdatedAt, nil,
		)
	suite.sqlmock.ExpectQuery(
		"SELECT * FROM `repayment` WHERE id = ? AND deleted_at IS NULL ORDER BY `repayment`.`id` LIMIT 1",
	).WithArgs(repayment.Id).WillReturnRows(repaymentRow)

	_, err := suite.repo.LoadRepayment(suite.ctx, repayment.Id)
	suite.Nil(err)
}
