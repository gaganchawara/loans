package repository

import (
	"context"
	"github.com/gaganchawara/loans/internal/loan/aggregate"

	"github.com/gaganchawara/loans/internal/loan/interfaces"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
	"gorm.io/gorm"
)

// repo is a struct representing the repository implementation.
type repo struct {
	db *gorm.DB
}

// NewRepository creates a new instance of the repository with the provided database connection.
func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repo{
		db: db,
	}
}

// CreateLoan creates a new loan record in the database.
func (r repo) CreateLoan(ctx context.Context, loan *entity.Loan) errors.Error {
	q := r.db.Table(loan.TableName()).Create(loan)
	if q.Error != nil {
		return errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
	}

	return nil
}

// UpdateLoan updates an existing loan record in the database.
func (r repo) UpdateLoan(ctx context.Context, loan *entity.Loan) errors.Error {
	q := r.db.Table(loan.TableName()).Updates(loan)
	if q.Error != nil {
		return errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
	}

	return nil
}

// LoadLoanAgg loads a loan aggregate, including the loan details and associated repayments.
func (r repo) LoadLoanAgg(ctx context.Context, loanId string) (*aggregate.LoanAgg, errors.Error) {
	loan, ierr := r.LoadLoan(ctx, loanId)
	if ierr != nil {
		return nil, ierr
	}

	repayments, ierr := r.LoadRepaymentsByLoanID(ctx, loanId)
	if ierr != nil {
		return nil, ierr
	}

	return &aggregate.LoanAgg{
		Loan:       loan,
		Repayments: repayments,
	}, nil
}

// LoadLoan loads a loan by its ID.
func (r repo) LoadLoan(ctx context.Context, loanId string) (*entity.Loan, errors.Error) {
	var loan entity.Loan
	q := r.db.Table(loan.TableName()).Where("id = ?", loanId).Where("deleted_at IS NULL").First(&loan)
	if q.Error != nil {
		if q.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(ctx, errorcode.NotFoundError, q.Error).Report()
		} else {
			return nil, errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
		}
	}

	return &loan, nil
}

// LoadRepaymentsByLoanID loads all repayments associated with a loan by loan ID.
func (r repo) LoadRepaymentsByLoanID(ctx context.Context, loanId string) ([]*entity.Repayment, errors.Error) {
	var repayments []*entity.Repayment
	q := r.db.Table(entity.TableRepayment).Where("loan_id = ?", loanId).Where("deleted_at IS NULL").
		Find(&repayments)
	if q.Error != nil {
		if q.Error == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
		}
	}

	return repayments, nil
}

// LoadRepayment loads a repayment by its ID.
func (r repo) LoadRepayment(ctx context.Context, id string) (*entity.Repayment, errors.Error) {
	var repayment entity.Repayment
	q := r.db.Table(repayment.TableName()).Where("id = ?", id).Where("deleted_at IS NULL").
		First(&repayment)
	if q.Error != nil {
		if q.Error == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
		}
	}

	return &repayment, nil
}
