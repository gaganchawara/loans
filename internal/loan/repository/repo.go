package repository

import (
	"context"

	"github.com/gaganchawara/loans/internal/loan/interfaces"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/internal/loan/entity"
	"github.com/gaganchawara/loans/pkg/errors"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repo{
		db: db,
	}
}

func (r repo) CreateLoan(ctx context.Context, loan *entity.Loan) errors.Error {
	q := r.db.Table(loan.TableName()).Create(loan)
	if q.Error != nil {
		return errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
	}

	return nil
}

func (r repo) UpdateLoan(ctx context.Context, loan *entity.Loan) errors.Error {
	q := r.db.Table(loan.TableName()).Updates(loan)
	if q.Error != nil {
		return errors.New(ctx, errorcode.InternalServerError, q.Error).Report()
	}

	return nil
}

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
